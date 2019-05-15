# Scaleway config

## TL;DR

Recommended config file:

```yaml
# get your access and secret keys on https://console.scaleway.com/account/credentials
access_key: SCWXXXXXXXXXXXXXXXXX
secret_key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
default_organization_id: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
default_region: fr-par
default_zone: fr-par-1
```

## Config file path

This package will try to locate the config file in the following ways:

1. Custom directory: `$SCW_CONFIG_PATH`
2. [XDG base directory](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html): `$XDG_CONFIG_HOME/scw/config.yaml`
3. Home directory: `$HOME/.config/scw/config.yaml` (`%USERPROFILE%/.config/scw/config.yaml` on windows)

## V1 config (DEPRECATED)

The V1 config `.scwrc` is supported but deprecated.
When found in the home directory, the V1 config is automatically migrated to a V2 config file in `$HOME/.config/scw/config.yaml`.
