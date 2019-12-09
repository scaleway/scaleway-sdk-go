# Changelog

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

