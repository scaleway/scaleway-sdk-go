# Changelog

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

