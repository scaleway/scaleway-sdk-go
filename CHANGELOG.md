# Changelog

## v1.0.0-beta.6 (2020-02-28)

### Features

* **account**: add projects ([#308](https://github.com/scaleway/scaleway-sdk-go/pull/308))
* **account**: remove projects ([#316](https://github.com/scaleway/scaleway-sdk-go/pull/316))
* **baremetal**: add zone in primary resource ([#305](https://github.com/scaleway/scaleway-sdk-go/pull/305))
* **baremetal**: get metrics ([#298](https://github.com/scaleway/scaleway-sdk-go/pull/298))
* **config**: add a list of tools that support this config file format ([#314](https://github.com/scaleway/scaleway-sdk-go/pull/314))
* **core**: add String method to scw.Money ([#284](https://github.com/scaleway/scaleway-sdk-go/pull/284))
* **core**: add send_usage setting in the config ([#273](https://github.com/scaleway/scaleway-sdk-go/pull/273))
* **core**: custom duration type ([#291](https://github.com/scaleway/scaleway-sdk-go/pull/291))
* **core**: handle instance quota exceeded error ([#287](https://github.com/scaleway/scaleway-sdk-go/pull/287))
* **core**: handle non standard errors ([#274](https://github.com/scaleway/scaleway-sdk-go/pull/274))
* **core**: rename send_usage into send_telemetry ([#313](https://github.com/scaleway/scaleway-sdk-go/pull/313))
* **core**: support more instance error types ([#278](https://github.com/scaleway/scaleway-sdk-go/pull/278))
* **domain**: add API ([#297](https://github.com/scaleway/scaleway-sdk-go/pull/297))
* **domain**: add record ViewConfig ([#317](https://github.com/scaleway/scaleway-sdk-go/pull/317))
* **errors**: Add ResourceExpiredError ([#280](https://github.com/scaleway/scaleway-sdk-go/pull/280))
* **httprecorder**: rename update cassette variable ([#300](https://github.com/scaleway/scaleway-sdk-go/pull/300))
* **instance**: add OrganizationDefault to UpdateSecurityGroup ([#279](https://github.com/scaleway/scaleway-sdk-go/pull/279))
* **instance**: add zone field in instance resources response ([#331](https://github.com/scaleway/scaleway-sdk-go/pull/331))
* **instance**: allow empty boot type on create server #325
* **instance**: use a zero IntervalStrategy when replaying cassettes ([#295](https://github.com/scaleway/scaleway-sdk-go/pull/295))
* **k8s**: add FeatureGates and AdmissionPlugins ([#289](https://github.com/scaleway/scaleway-sdk-go/pull/289))
* **k8s**: add WaitForPoolNodesReady & WaitForClusterNodesReady helper methods ([#312](https://github.com/scaleway/scaleway-sdk-go/pull/312))
* **k8s**: add creation_error node status ([#328](https://github.com/scaleway/scaleway-sdk-go/pull/328))
* **k8s**: add upgrading pool status ([#319](https://github.com/scaleway/scaleway-sdk-go/pull/319))
* **lb**: add CreateIP ([#290](https://github.com/scaleway/scaleway-sdk-go/pull/290))
* **lb**: proxy protocol ([#299](https://github.com/scaleway/scaleway-sdk-go/pull/299))
* **rdb**: add ExportDatabaseBackup ([#292](https://github.com/scaleway/scaleway-sdk-go/pull/292))
* **rdb**: add OrganizationID field to ListDatabaseBackups ([#321](https://github.com/scaleway/scaleway-sdk-go/pull/321))
* **scw**: GetCacheDirectory ([#304](https://github.com/scaleway/scaleway-sdk-go/pull/304))
* **strcase**: add functions ([#275](https://github.com/scaleway/scaleway-sdk-go/pull/275))
* **test**: add human name ([#309](https://github.com/scaleway/scaleway-sdk-go/pull/309))
* update generated apis ([#276](https://github.com/scaleway/scaleway-sdk-go/pull/276))
* update generated apis ([#285](https://github.com/scaleway/scaleway-sdk-go/pull/285))

### Fixes

* **core**: Size typo ([#293](https://github.com/scaleway/scaleway-sdk-go/pull/293))
* **core**: do not omit empty Money fields ([#288](https://github.com/scaleway/scaleway-sdk-go/pull/288))
* **core**: handle content-types in ResponseError ([#315](https://github.com/scaleway/scaleway-sdk-go/pull/315))
* **core**: handle precision and clean rounding of Money ([#286](https://github.com/scaleway/scaleway-sdk-go/pull/286))
* **instance**: set all server user data stop deleting all keys ([#281](https://github.com/scaleway/scaleway-sdk-go/pull/281))
* **k8s**: add default timeout to WaitForCluster ([#323](https://github.com/scaleway/scaleway-sdk-go/pull/323))
* **k8s**: copy helpers from v1beta3 to beta4 ([#277](https://github.com/scaleway/scaleway-sdk-go/pull/277))
* **k8s**: make WaitForPool coherent with others ([#333](https://github.com/scaleway/scaleway-sdk-go/pull/333))
* **k8s**: remove WaitForClusterPools ([#334](https://github.com/scaleway/scaleway-sdk-go/pull/334))
* **lb**: UpdateIP method reverse field ([#320](https://github.com/scaleway/scaleway-sdk-go/pull/320))
* **scripts**: golangci-lint binary installation test ([#301](https://github.com/scaleway/scaleway-sdk-go/pull/301))
* **scw**: money type now implement stringer without pointer ([#303](https://github.com/scaleway/scaleway-sdk-go/pull/303))
* properly convert ipId to ip-id, rename field of DeleteIpRequest ([#272](https://github.com/scaleway/scaleway-sdk-go/pull/272))

### Others

* **chore - core**: add windows and macos build tests ([#336](https://github.com/scaleway/scaleway-sdk-go/pull/336))
* **chore**: add github action for PR testing ([#332](https://github.com/scaleway/scaleway-sdk-go/pull/332))
* **chore**: add linters ([#294](https://github.com/scaleway/scaleway-sdk-go/pull/294))
* **chore**: post release commit ([#271](https://github.com/scaleway/scaleway-sdk-go/pull/271))
* **chore**: remove go 1.10 and 1.11 from gh action ([#335](https://github.com/scaleway/scaleway-sdk-go/pull/335))
* **chore**: replace AttachIp
* **chore**: untitle first word of a string ([#338](https://github.com/scaleway/scaleway-sdk-go/pull/338))
* **chore**: update release script ([#326](https://github.com/scaleway/scaleway-sdk-go/pull/326))
* **doc - instance**: improve create image doc ([#324](https://github.com/scaleway/scaleway-sdk-go/pull/324))
* **doc - instance**: improve documentation ([#318](https://github.com/scaleway/scaleway-sdk-go/pull/318))
* **doc**: add documentation for namespace naming ([#302](https://github.com/scaleway/scaleway-sdk-go/pull/302))
* **doc**: improve instance images documentation ([#322](https://github.com/scaleway/scaleway-sdk-go/pull/322))
* **refactor - instance**: update setSecurityGroupRequest fields order



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

