# Changelog

## v1.0.0-beta.7 (2020-10-07)

### Features

* **account**: add project id to ssh key object ([#447](https://github.com/scaleway/scaleway-sdk-go/pull/447))
* **account**: add project_id in ssh-key ([#438](https://github.com/scaleway/scaleway-sdk-go/pull/438))

* **baremetal**: add RetryInterval variable ([#357](https://github.com/scaleway/scaleway-sdk-go/pull/357))
* **baremetal**: add boot type in start server ([#500](https://github.com/scaleway/scaleway-sdk-go/pull/500))
* **baremetal**: add install in create server ([#499](https://github.com/scaleway/scaleway-sdk-go/pull/499))
* **baremetal**: add ping status in server ([#390](https://github.com/scaleway/scaleway-sdk-go/pull/390))
* **baremetal**: add project-id to ipfailover v1alpha1 ([#559](https://github.com/scaleway/scaleway-sdk-go/pull/559))
* **baremetal**: add server helpers func in v1 ([#397](https://github.com/scaleway/scaleway-sdk-go/pull/397))
* **baremetal**: add support for projects ([#535](https://github.com/scaleway/scaleway-sdk-go/pull/535))
* **baremetal**: add v1 ([#396](https://github.com/scaleway/scaleway-sdk-go/pull/396))

* **config**: add support for default profile ([#478](https://github.com/scaleway/scaleway-sdk-go/pull/478))
* **config**: move telemetry config in profile ([#427](https://github.com/scaleway/scaleway-sdk-go/pull/427))

* **core**: add GetSecretKey and GetAccessKey ([#386](https://github.com/scaleway/scaleway-sdk-go/pull/386))

* **domain**: add domain registration process on DomainSummary ([#560](https://github.com/scaleway/scaleway-sdk-go/pull/560))
* **domain**: add external domain support ([#401](https://github.com/scaleway/scaleway-sdk-go/pull/401))
* **domain**: add new task types ([#415](https://github.com/scaleway/scaleway-sdk-go/pull/415))
* **domain**: add new types in messages and email validated in domain contact ([#423](https://github.com/scaleway/scaleway-sdk-go/pull/423))
* **domain**: add organization_id filter ([#424](https://github.com/scaleway/scaleway-sdk-go/pull/424))
* **domain**: add registration process and is_external filter ([#507](https://github.com/scaleway/scaleway-sdk-go/pull/507))

* **error**: add resource locked error ([#361](https://github.com/scaleway/scaleway-sdk-go/pull/361))
* **errors**: handle instance unknown resource ([#374](https://github.com/scaleway/scaleway-sdk-go/pull/374))

* **instance**: Add private_network filter on ListServers ([#567](https://github.com/scaleway/scaleway-sdk-go/pull/567))
* **instance**: Remove private_networks in CreateServerRequest ([#572](https://github.com/scaleway/scaleway-sdk-go/pull/572))
* **instance**: add a WaitForImage method ([#381](https://github.com/scaleway/scaleway-sdk-go/pull/381))
* **instance**: add a wait on volume ([#418](https://github.com/scaleway/scaleway-sdk-go/pull/418))
* **instance**: add project to others resources ([#519](https://github.com/scaleway/scaleway-sdk-go/pull/519))
* **instance**: add project to resource IP ([#460](https://github.com/scaleway/scaleway-sdk-go/pull/460))
* **instance**: add snapshot wait utils ([#398](https://github.com/scaleway/scaleway-sdk-go/pull/398))
* **instance**: add support for name in backup action ([#383](https://github.com/scaleway/scaleway-sdk-go/pull/383))
* **instance**: add tags to IP ([#344](https://github.com/scaleway/scaleway-sdk-go/pull/344))
* **instance**: add volume type endpoint ([#440](https://github.com/scaleway/scaleway-sdk-go/pull/440))
* **instance**: allow filter by tags in server list ([#373](https://github.com/scaleway/scaleway-sdk-go/pull/373))
* **instance**: exclude SetIp of the SDK ([#461](https://github.com/scaleway/scaleway-sdk-go/pull/461))
* **instance**: fix volume template project oneof ([#561](https://github.com/scaleway/scaleway-sdk-go/pull/561))
* **instance**: fix volume type endpoint ([#441](https://github.com/scaleway/scaleway-sdk-go/pull/441))
* **instance**: update metadata struct ([#541](https://github.com/scaleway/scaleway-sdk-go/pull/541))
* **instance**: use generated UpdateVolume ([#417](https://github.com/scaleway/scaleway-sdk-go/pull/417))

* **iot**: add REST network type ([#505](https://github.com/scaleway/scaleway-sdk-go/pull/505))
* **iot**: add WaitForHub IoT helper ([#445](https://github.com/scaleway/scaleway-sdk-go/pull/445))
* **iot**: add hub events settings ([#476](https://github.com/scaleway/scaleway-sdk-go/pull/476))
* **iot**: add name generation on hub network and device ([#526](https://github.com/scaleway/scaleway-sdk-go/pull/526))
* **iot**: add product plan in UpdateHubRequest ([#513](https://github.com/scaleway/scaleway-sdk-go/pull/513))
* **iot**: add support for hub-id in an UpdateDeviceRequest #554
* **iot**: enable sdk generation ([#444](https://github.com/scaleway/scaleway-sdk-go/pull/444))

* **k8s**: add kubeconfig and helpers ([#364](https://github.com/scaleway/scaleway-sdk-go/pull/364))
* **k8s**: add project support ([#517](https://github.com/scaleway/scaleway-sdk-go/pull/517))
* **k8s**: add scaledown unneeded time ([#385](https://github.com/scaleway/scaleway-sdk-go/pull/385))
* **k8s**: add traefik2 ingress ([#443](https://github.com/scaleway/scaleway-sdk-go/pull/443))
* **k8s**: add v1 api
* **k8s**: add wait for node method ([#352](https://github.com/scaleway/scaleway-sdk-go/pull/352))
* **k8s**: flag to delete block and pvc with kapsule ([#416](https://github.com/scaleway/scaleway-sdk-go/pull/416))

* **lb**: add SSL compatibility level ([#343](https://github.com/scaleway/scaleway-sdk-go/pull/343))
* **lb**: multi-certificacte in frontend ([#492](https://github.com/scaleway/scaleway-sdk-go/pull/492))
* **lb**: add default on the lb name generation ([#493](https://github.com/scaleway/scaleway-sdk-go/pull/493))
* **lb**: add private network ([#402](https://github.com/scaleway/scaleway-sdk-go/pull/402))
* **lb**: add support for multiple certificates in frontend ([#479](https://github.com/scaleway/scaleway-sdk-go/pull/479))
* **lb**: remove subscriber from cli generation ([#490](https://github.com/scaleway/scaleway-sdk-go/pull/490))
* **lb**: support projects ([#504](https://github.com/scaleway/scaleway-sdk-go/pull/504))

* **rdb**: add GetInstanceLogs method ([#484](https://github.com/scaleway/scaleway-sdk-go/pull/484))
* **rdb**: add beta flags ([#411](https://github.com/scaleway/scaleway-sdk-go/pull/411))
* **rdb**: add generated prefix ([#558](https://github.com/scaleway/scaleway-sdk-go/pull/558))
* **rdb**: add project_id to resources ([#571](https://github.com/scaleway/scaleway-sdk-go/pull/571))
* **rdb**: allow setting initial settings while creating an RDB instance. ([#536](https://github.com/scaleway/scaleway-sdk-go/pull/536))
* **rdb**: update upgrade call to add non-ha to ha upgrade ([#451](https://github.com/scaleway/scaleway-sdk-go/pull/451))

* **registry**: add wait functions on image and tag ([#426](https://github.com/scaleway/scaleway-sdk-go/pull/426))
* **registry**: support projects ([#506](https://github.com/scaleway/scaleway-sdk-go/pull/506))

* **vpc**: add support for project ([#565](https://github.com/scaleway/scaleway-sdk-go/pull/565))
* **vpc**: generate doc/cli/sdk ([#532](https://github.com/scaleway/scaleway-sdk-go/pull/532))

* add configurable retryInterval and timeout ([#428](https://github.com/scaleway/scaleway-sdk-go/pull/428))
* add projects ([#452](https://github.com/scaleway/scaleway-sdk-go/pull/452))
* add support for pl-waw-1 zone and pl-waw region ([#557](https://github.com/scaleway/scaleway-sdk-go/pull/557))
* set default project client value when request is empty ([#457](https://github.com/scaleway/scaleway-sdk-go/pull/457))
* use pointer to time.Time to allow null value ([#523](https://github.com/scaleway/scaleway-sdk-go/pull/523))

### Fixes

* **config**: LoadConfigFromPath better handle cross platform error ([#395](https://github.com/scaleway/scaleway-sdk-go/pull/395))
* **config**: debug level when reading env ([#349](https://github.com/scaleway/scaleway-sdk-go/pull/349))
* **errors**: handle all instance not found return messages ([#432](https://github.com/scaleway/scaleway-sdk-go/pull/432))
* **instance**: GetServerTypesAvailabilityResponse nested object
* **instance**: volume listing return all types by default
* **k8s**: add global retry interval for wait func ([#354](https://github.com/scaleway/scaleway-sdk-go/pull/354))
* **k8s**: change type for kubeconfig certificates ([#362](https://github.com/scaleway/scaleway-sdk-go/pull/362))
* **k8s**: panic on helpers timeout ([#369](https://github.com/scaleway/scaleway-sdk-go/pull/369))
* **k8s**: remove oldbinpacking from autoscaler estimator ([#389](https://github.com/scaleway/scaleway-sdk-go/pull/389))
* **rdb**: switch from ip to ipnet in ACL ([#482](https://github.com/scaleway/scaleway-sdk-go/pull/482))
* **registry**: use scw.Size ([#391](https://github.com/scaleway/scaleway-sdk-go/pull/391))
* date format in query parameters ([#471](https://github.com/scaleway/scaleway-sdk-go/pull/471))
* handle additional non standards errors ([#525](https://github.com/scaleway/scaleway-sdk-go/pull/525))

### Others

* **chore - baremetal**: rename to GetOfferFromOfferNameRequest ([#353](https://github.com/scaleway/scaleway-sdk-go/pull/353))
* **chore - config**: use os.UserHomeDir
* **chore - iot**: deprecate organization ID in CreateNetworkRequest ([#501](https://github.com/scaleway/scaleway-sdk-go/pull/501))
* **chore - locality**: small fixes ([#455](https://github.com/scaleway/scaleway-sdk-go/pull/455))
* **chore - rdb**: add locked status ([#568](https://github.com/scaleway/scaleway-sdk-go/pull/568))
* **chore - sdk_compilation_test**: add missing api packages in test ([#446](https://github.com/scaleway/scaleway-sdk-go/pull/446))
* **chore - vendor**: remove vendor folder and rely on go module instead ([#469](https://github.com/scaleway/scaleway-sdk-go/pull/469))
* **chore**: export Scaleway environment variable constants ([#400](https://github.com/scaleway/scaleway-sdk-go/pull/400))
* **chore**: export scw.getScwConfigDir ([#514](https://github.com/scaleway/scaleway-sdk-go/pull/514))
* **chore**: use focal fossa as an example ([#408](https://github.com/scaleway/scaleway-sdk-go/pull/408))
* **doc - account**: improve ssh-key documentation ([#375](https://github.com/scaleway/scaleway-sdk-go/pull/375))
* **doc - config**: add project_id ([#543](https://github.com/scaleway/scaleway-sdk-go/pull/543))
* **doc - core**: update credentials links ([#564](https://github.com/scaleway/scaleway-sdk-go/pull/564))
* **doc - instance**: documentation improvments ([#563](https://github.com/scaleway/scaleway-sdk-go/pull/563))
* **doc - instance**: fix project documentation ([#551](https://github.com/scaleway/scaleway-sdk-go/pull/551))
* **doc - rdb**: add go doc ([#465](https://github.com/scaleway/scaleway-sdk-go/pull/465))
* **doc - rdb**: remove dev1-s mention ([#472](https://github.com/scaleway/scaleway-sdk-go/pull/472))
* **docs**: add deprecation notices ([#570](https://github.com/scaleway/scaleway-sdk-go/pull/570))
* **docs**: improve API documentations ([#345](https://github.com/scaleway/scaleway-sdk-go/pull/345))
* **docs**: update generated godoc ([#481](https://github.com/scaleway/scaleway-sdk-go/pull/481))
* **docs - account**: comment few fields #380
* **docs - baremetal**: update comment ([#503](https://github.com/scaleway/scaleway-sdk-go/pull/503))
* **docs**: fix badges and update godoc links ([#520](https://github.com/scaleway/scaleway-sdk-go/pull/520))
* **rdb**: add support for wait on backup resource ([#474](https://github.com/scaleway/scaleway-sdk-go/pull/474))
* **rdb**: add wait function for instance logs ([#485](https://github.com/scaleway/scaleway-sdk-go/pull/485))
* **refactor**: rename method DurationPtr



## v1.0.0-beta.6 (2020-02-28)

### Features

* **baremetal**: add zone in primary resource ([#305](https://github.com/scaleway/scaleway-sdk-go/pull/305))
* **baremetal**: add get metrics method ([#298](https://github.com/scaleway/scaleway-sdk-go/pull/298))
* **core**: add String method to scw.Money ([#284](https://github.com/scaleway/scaleway-sdk-go/pull/284))
* **core**: add send_telemetry setting in the config ([#273](https://github.com/scaleway/scaleway-sdk-go/pull/273))
* **core**: add custom duration type ([#291](https://github.com/scaleway/scaleway-sdk-go/pull/291))
* **core**: handle instance quota exceeded error ([#287](https://github.com/scaleway/scaleway-sdk-go/pull/287))
* **core**: handle non standard errors ([#274](https://github.com/scaleway/scaleway-sdk-go/pull/274))
* **core**: support more instance error types ([#278](https://github.com/scaleway/scaleway-sdk-go/pull/278))
* **core**: Add ResourceExpiredError ([#280](https://github.com/scaleway/scaleway-sdk-go/pull/280))
* **core**: add GetCacheDirectory method ([#304](https://github.com/scaleway/scaleway-sdk-go/pull/304))
* **domain**: first release of the API ([#297](https://github.com/scaleway/scaleway-sdk-go/pull/297))
* **instance**: add OrganizationDefault to UpdateSecurityGroup ([#279](https://github.com/scaleway/scaleway-sdk-go/pull/279))
* **instance**: add zone field in instance resources response ([#331](https://github.com/scaleway/scaleway-sdk-go/pull/331))
* **instance**: allow empty boot type on create server [#325](https://github.com/scaleway/scaleway-sdk-go/pull/325))
* **k8s**: add FeatureGates and AdmissionPlugins ([#289](https://github.com/scaleway/scaleway-sdk-go/pull/289))
* **k8s**: add WaitForPoolNodesReady and WaitForClusterNodesReady helper methods ([#312](https://github.com/scaleway/scaleway-sdk-go/pull/312))
* **k8s**: add creation_error node status ([#328](https://github.com/scaleway/scaleway-sdk-go/pull/328))
* **k8s**: add upgrading pool status ([#319](https://github.com/scaleway/scaleway-sdk-go/pull/319))
* **k8s**: add UpgradeAvailable in pool ([#276](https://github.com/scaleway/scaleway-sdk-go/pull/276))
* **k8s**: add a field to rename Cluster ([#285](https://github.com/scaleway/scaleway-sdk-go/pull/285))
* **k8s**: add v1beta4 version ([#276](https://github.com/scaleway/scaleway-sdk-go/pull/276))
* **lb**: add CreateIP method ([#290](https://github.com/scaleway/scaleway-sdk-go/pull/290))
* **lb**: add support for proxy protocol ([#299](https://github.com/scaleway/scaleway-sdk-go/pull/299))
* **lb**: add Subscriber related methods ([#276](https://github.com/scaleway/scaleway-sdk-go/pull/276))
* **rdb**: add ExportDatabaseBackup method ([#292](https://github.com/scaleway/scaleway-sdk-go/pull/292))
* **rdb**: add OrganizationID field to ListDatabaseBackups ([#321](https://github.com/scaleway/scaleway-sdk-go/pull/321))
* **strcase**: add functions ([#275](https://github.com/scaleway/scaleway-sdk-go/pull/275))

### Fixes

* **core**: doc typo ([#293](https://github.com/scaleway/scaleway-sdk-go/pull/293))
* **core**: do not omit empty Money fields ([#288](https://github.com/scaleway/scaleway-sdk-go/pull/288))
* **core**: handle content-types in ResponseError ([#315](https://github.com/scaleway/scaleway-sdk-go/pull/315))
* **core**: handle precision and clean rounding of Money ([#286](https://github.com/scaleway/scaleway-sdk-go/pull/286))
* **instance**: set all server user data stop deleting all keys ([#281](https://github.com/scaleway/scaleway-sdk-go/pull/281))
* **instance**: properly convert ipId to ip-id, rename field of DeleteIpRequest ([#272](https://github.com/scaleway/scaleway-sdk-go/pull/272))
* **lb**: UpdateIP method reverse field ([#320](https://github.com/scaleway/scaleway-sdk-go/pull/320))
* **scripts**: golangci-lint binary installation test ([#301](https://github.com/scaleway/scaleway-sdk-go/pull/301))


## v1.0.0-beta.5 (2019-12-09)

### Features

* **baremetal**: add ListOffers and ListOs methods ([#259](https://github.com/scaleway/scaleway-sdk-go/pull/259))
* **config**: add a typed error on config file not found ([#264](https://github.com/scaleway/scaleway-sdk-go/pull/264))
* **config**: add a commented configuration file ([#231](https://github.com/scaleway/scaleway-sdk-go/pull/231))
* **core**: get Region from Zone ([#255](https://github.com/scaleway/scaleway-sdk-go/pull/255))
* **core**: client validation ([#238](https://github.com/scaleway/scaleway-sdk-go/pull/238))
* **core**: IPNet type ([#236](https://github.com/scaleway/scaleway-sdk-go/pull/236))
* **core**: introduce format validation in locality parsing ([#237](https://github.com/scaleway/scaleway-sdk-go/pull/237))
* **core**: add scw.MergeProfiles command ([#234](https://github.com/scaleway/scaleway-sdk-go/pull/234))
* **core**: add ClientCredentialError ([#228](https://github.com/scaleway/scaleway-sdk-go/pull/228))
* **core**: client without auth by default ([#233](https://github.com/scaleway/scaleway-sdk-go/pull/233))
* **k8s**: add maintenance and upgrade features ([#258](https://github.com/scaleway/scaleway-sdk-go/pull/258))
* **k8s**: add Pool.PlacementGroupID field ([#246](https://github.com/scaleway/scaleway-sdk-go/pull/246))
* **lb**: custom certificate ([#262](https://github.com/scaleway/scaleway-sdk-go/pull/262))
* **lb**: add ListBackendStats method ([#252](https://github.com/scaleway/scaleway-sdk-go/pull/252))
* **instance**: publish WaitForServer ([#244](https://github.com/scaleway/scaleway-sdk-go/pull/244))
* **instance**: generate name for new snapshot or image ([#230](https://github.com/scaleway/scaleway-sdk-go/pull/230))
* **strcase**: add strcase lib ([#229](https://github.com/scaleway/scaleway-sdk-go/pull/229))
* **rdb**: wait for Instance ([#249](https://github.com/scaleway/scaleway-sdk-go/pull/249))
* **rdb**: add rdb API ([#247](https://github.com/scaleway/scaleway-sdk-go/pull/247))
* **registry**: add WaitForNamespace method ([#253](https://github.com/scaleway/scaleway-sdk-go/pull/253))
* **validation**: add IsEmail ([#242](https://github.com/scaleway/scaleway-sdk-go/pull/242))
* **validation**: make validation package public ([#241](https://github.com/scaleway/scaleway-sdk-go/pull/241))

### Fixes

* **baremetal**: fix WaitForServer and add WaitForServerInstall ([#263](https://github.com/scaleway/scaleway-sdk-go/pull/263))
* **baremetal**: add Undelivered and Locked to terminalStatus ([#260](https://github.com/scaleway/scaleway-sdk-go/pull/260))
* **config**: merge selected profile on top of default profile ([#243](https://github.com/scaleway/scaleway-sdk-go/pull/243))
* **instance**: WaitForServer returns an error interface ([#245](https://github.com/scaleway/scaleway-sdk-go/pull/245))
* **instance**: use IPNet type for security group rule ip_range ([#240](https://github.com/scaleway/scaleway-sdk-go/pull/240))
* **instance**: update placement-group now works ([#224](https://github.com/scaleway/scaleway-sdk-go/pull/224))


## v1.0.0-beta.4 (2019-10-25)

### Breaking Changes

* use uint32 for page_size ([#210](https://github.com/scaleway/scaleway-sdk-go/pull/210))

### Features

* update generated apis ([#218](https://github.com/scaleway/scaleway-sdk-go/pull/218))
* update generated apis ([#216](https://github.com/scaleway/scaleway-sdk-go/pull/216))
* **lb**: add WaitForLb method ([#212](https://github.com/scaleway/scaleway-sdk-go/pull/212))
* update generated apis ([#213](https://github.com/scaleway/scaleway-sdk-go/pull/213))
* update generated apis ([#208](https://github.com/scaleway/scaleway-sdk-go/pull/208))
* **marketplace**: uppercase commercial type in GetLocalImageIDByLabel ([#205](https://github.com/scaleway/scaleway-sdk-go/pull/205))
* add kubeconfig helpers for k8s ([#204](https://github.com/scaleway/scaleway-sdk-go/pull/204))
* update generated apis ([#203](https://github.com/scaleway/scaleway-sdk-go/pull/203))
* add k8s WaitForCluster method ([#202](https://github.com/scaleway/scaleway-sdk-go/pull/202))
* add scw.IPPtr ([#200](https://github.com/scaleway/scaleway-sdk-go/pull/200))
* add k8s v1beta3 ([#198](https://github.com/scaleway/scaleway-sdk-go/pull/198))

### Fixes

* **instance**: update generated apis ([#219](https://github.com/scaleway/scaleway-sdk-go/pull/219))
* **instance**: ListImage total count ([#209](https://github.com/scaleway/scaleway-sdk-go/pull/209))
* cleanup unused code ([#217](https://github.com/scaleway/scaleway-sdk-go/pull/217))
* **scw.File**: add unmarshal ([#201](https://github.com/scaleway/scaleway-sdk-go/pull/201))

### Documentation

* fix examples ([#215](https://github.com/scaleway/scaleway-sdk-go/pull/215))

## v1.0.0-beta.3 (2019-10-01)

### Features

* add interface body getter to std err ([#192](https://github.com/scaleway/scaleway-sdk-go/pull/192))
* add support for out of stock error ([#190](https://github.com/scaleway/scaleway-sdk-go/pull/190))
* use uint32 for page count ([#193](https://github.com/scaleway/scaleway-sdk-go/pull/193))
* add raw body to standard errors ([#191](https://github.com/scaleway/scaleway-sdk-go/pull/191))
* update generated apis ([#188](https://github.com/scaleway/scaleway-sdk-go/pull/188))

### Fixes

* e2e tests project ([#195](https://github.com/scaleway/scaleway-sdk-go/pull/195))
* **instance**: allow image to be empty in CreateServer ([#189](https://github.com/scaleway/scaleway-sdk-go/pull/189))


## v1.0.0-beta.2 (2019-09-16)

### Features

* test standard errors ([#183](https://github.com/scaleway/scaleway-sdk-go/pull/183))
* update generated apis ([#182](https://github.com/scaleway/scaleway-sdk-go/pull/182))
* standard error structures ([#177](https://github.com/scaleway/scaleway-sdk-go/pull/177))

### Fixes

* attach volume key choice ([#184](https://github.com/scaleway/scaleway-sdk-go/pull/184))
* add user agent to e2e tests ([#181](https://github.com/scaleway/scaleway-sdk-go/pull/181))
* update version ([#180](https://github.com/scaleway/scaleway-sdk-go/pull/180))
* UpdateSecurityGroupRule can now remove DestPortFrom and DestPortTo ([#179](https://github.com/scaleway/scaleway-sdk-go/pull/179))

