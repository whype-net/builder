# Builder

Small CLI utillity to query SteamDB for the Rust Dedicated Server depot information

## Usage

```sh
$ builder check
2023/10/23 00:28:26 'public' branch last updated at 1697442502.
2023/10/23 00:28:26 Build ID: 12445585
2023/10/23 00:28:26 Linx Manifest ID: 6107145083830036397
2023/10/23 00:28:26 Common Manifest ID: 6661281094410849912"
```

**Export results**

Places results in a `source`able format to aid building

```sh
$ builder check --export >> build.env
```
