const { spawnSync } = require("child_process"),
  readline = require("readline"),
  fs = require("fs")
;

const gitRawCommits = require("git-raw-commits"),
  semver = require("semver"),
  getStream = require("get-stream"),
  _ = require("colors")
;

const _typeReg = /(?<type>[a-zA-Z]+)/;
const _scopeReg = /(\((?<scope>.*)\))?/;
const _messageReg = /(?<message>[^(]*)/;
const _mrReg = /(\(#(?<mr>[0-9]+)\))?/;
const COMMIT_REGEX = new RegExp(`${_typeReg.source}${_scopeReg.source}: *${_messageReg.source} *${_mrReg.source}`);
const CHANGELOG_PATH = "../../CHANGELOG.md";
const GO_VERSION_PATH = "../../scw/version.go";
const TMP_BRANCH = "new-release";
const TMP_REMOTE = "scaleway-release";
const REPO_URL = "git@github.com:scaleway/scaleway-sdk-go.git";

async function main() {

  //
  // Initialization
  //
  console.log("Adding temporary remote on local repo".blue);
  git( "remote", "add", TMP_REMOTE, REPO_URL);
  console.log(`   Successfully created ${TMP_REMOTE} remote`.green);

  console.log("Make sure we are working on an up to date master".blue);
  git( "fetch", TMP_REMOTE);
  git( "checkout", `${TMP_REMOTE}/master`);
  console.log(`   Successfully created ${TMP_REMOTE} remote`.green);

  console.log("Trying to find last release tag".blue);
  const lastSemverTag = git("tag")
    .trim()
    .split("\n")
    .filter(semver.valid)
    .sort((a, b) => semver.rcompare(semver.clean(a), semver.clean(b)))[0];
  console.log(`    Last found release tag was ${lastSemverTag}`.green);

  console.log("Listing commit since last release".blue);
  const commits = (await getStream.array(gitRawCommits({ from: lastSemverTag, format: "%s" }))).map(c => c.toString().trim());
  commits.forEach(c => console.log(`    ${c}`.grey));
  console.log(`    We found ${commits.length} commits since last release`.green);

  const newVersion = semver.clean(await prompt("Enter new version: ".magenta));
  if (!newVersion) {
    throw new Error(`invalid version`);
  }

  //
  // Creating release commit
  //

  console.log(`Updating ${CHANGELOG_PATH} and ${GO_VERSION_PATH}`.blue);
  // Generate changelog lines by section
  const changelogLines = { feat: [], fix: [] };
  commits.forEach(commit => {
    const result = COMMIT_REGEX.exec(commit);
    if (!result || !(result.groups.type in changelogLines)) {
      console.warn(`WARNING: Ignoring commit ${commit}`.yellow);
      return;
    }
    const stdCommit = result.groups;

    let line = [`*`, stdCommit.scope ? `**${stdCommit.scope}**:` : "", stdCommit.message, stdCommit.mr ? `([#${stdCommit.mr}](https://github.com/scaleway/scaleway-sdk-go/pull/${stdCommit.mr}))` : ""]
      .map(s => s.trim())
      .filter(v => v)
      .join(" ");
    changelogLines[stdCommit.type].push(line);
  });

  const changelogHeader = `## v${newVersion} (${new Date().toISOString().substring(0, 10)})`;
  const changelogSections = [];
  if (changelogLines.feat) {
    changelogSections.push("### Features\n\n" + changelogLines.feat.join("\n"));
  }
  if (changelogLines.fix) {
    changelogSections.push("### Fixes\n\n" + changelogLines.fix.join("\n"));
  }
  const changelogBody = changelogSections.join("\n\n");
  const changelog = `${changelogHeader}\n\n${changelogBody}`;

  replaceInFile(CHANGELOG_PATH, "# Changelog", "# Changelog\n\n" + changelog + "\n");
  replaceInFile(GO_VERSION_PATH, /const version[^\n]*\n/, `const version = "v${newVersion}"\n`);
  console.log(`    Update success`.green);

  await prompt(`Please review ${CHANGELOG_PATH} and ${GO_VERSION_PATH}. When everything is fine hit enter to continue ...`.magenta);

  console.log(`Creating release commit`.blue);
  git("checkout", "-b", TMP_BRANCH);
  git("add", CHANGELOG_PATH, GO_VERSION_PATH);
  git("commit", "-m", `chore: release ${newVersion}`);
  git("push", "-f", "--set-upstream", TMP_REMOTE, TMP_BRANCH);

  await prompt(`Please create an PR here: https://github.com/scaleway/scaleway-sdk-go/pull/new/new-release . Hit enter when its merged .....`.magenta);

  console.log("Time to create a github release with the following info\n".blue);
  console.log(`Title: v${newVersion}\n\n`.gray);
  console.log(`${changelogBody}\n\n`.gray);
  await prompt(`You should create a new github release here: https://github.com/scaleway/scaleway-sdk-go/releases/new/ . Hit enter when the new release is created .....`.magenta);

  //
  // Creating post release commit
  //
  console.log("Make sure we pull the latest commit from master".blue);
  git("fetch", TMP_REMOTE);
  git("checkout", `${TMP_REMOTE}/master`);
  console.log("    Successfully checkout upstream/master".green);

  console.log(`Creating post release commit`.blue);
  git("branch", "-D", TMP_BRANCH);
  git("checkout", "-b", TMP_BRANCH);
  replaceInFile(GO_VERSION_PATH, /const version[^\n]*\n/, `const version = "v${newVersion}+dev"\n`);
  git("add", GO_VERSION_PATH);
  git("commit", "-m", "chore: post release commit");
  git("push", "-f", "--set-upstream", TMP_REMOTE, TMP_BRANCH);
  git("checkout", "master");
  git("branch", "-D", TMP_BRANCH);
  await prompt(`Please create an PR here: https://github.com/scaleway/scaleway-sdk-go/pull/new/new-release . Hit enter when its merged .....`.magenta);

  console.log("Make sure we pull the latest commit from master".blue);
  git("pull", TMP_REMOTE, "master");
  console.log("    Successfully pull master".green);

  console.log("Remove temporary remote".blue);
  git("remote", "remove", TMP_REMOTE);
  console.log("    Successfully remove temporary remote".green);

  console.log(`🚀 Release Success `.green);
}

function git(...args) {
  console.log(`    git ${args.join(" ")}`.grey);
  const { stdout, status, stderr } = spawnSync("git", args, { encoding: "utf8" });
  if (status !== 0) {
    throw new Error(`return status ${status}\n${stderr}\n`);
  }
  return stdout;
}

function replaceInFile(path, oldStr, newStr) {
  console.log(`    Editing ${path}`.grey);
  const content = fs.readFileSync(path, { encoding: "utf8" }).replace(oldStr, newStr);
  fs.writeFileSync(path, content);
}

function prompt(prompt) {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });
  return new Promise(resolve => {
    rl.question(prompt, answer => {
      resolve(answer);
      rl.close();
    });
  });
}

main().catch(console.error);
