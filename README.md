# Mateus' F-Droid repo

A personal, unofficial [F-Droid](https://f-droid.org/) repository. It mirrors the
original, developer-signed release APKs of a few apps so they can be installed and
kept up to date through an F-Droid client.

Only this repo's **index** is signed by me. Each APK keeps its **original
developer signature**, and the index pins each package to that developer's
signing key, so a compromise of this repo or its host cannot push a tampered
"update".

### Apps

<!-- This table is auto-generated. Do not edit -->
| Icon | Name | Description | Version |
| --- | --- | --- | --- |
| | [**Continuum**](https://github.com/cygnusx-1-org/continuum) | Ad-free Reddit client for Android (fork of Infinity for Reddit) | _pending first CI run_ |
<!-- end apps table -->

### How to use

1. Install an F-Droid client — [F-Droid](https://f-droid.org/) or a fork like
   [Droid-ify](https://github.com/Droid-ify/client).
2. Add this repository:

   ```
   https://raw.githubusercontent.com/mateus2k2/fdroid/main/fdroid/repo?fingerprint=d5c9373ac0431150dc54b59a6364e1dfb0ead9ef61591e14f95845168eb7a794
   ```

   Or scan this QR code:

   <p align="center">
     <img src=".github/qrcode.png?raw=true" alt="F-Droid repo QR code"/>
   </p>

3. Open the link in your F-Droid client and confirm the repository (URL and
   fingerprint are pre-filled).
4. Search for the app (e.g. "Continuum") and install it. Updates then arrive
   automatically.

### How it works

A scheduled GitHub Action polls each app's upstream GitHub releases, downloads
new release APKs, rebuilds and signs the repo index with
[`fdroidserver`](https://gitlab.com/fdroid/fdroidserver), and commits the result
to `fdroid/repo/`, which is served raw from this repo. See [setup.md](setup.md)
for how the machinery is wired, and [SETUP-CONTINUUM.md](SETUP-CONTINUUM.md) for
this repo's specifics.

Built on the [xarantolus/fdroid](https://github.com/xarantolus/fdroid) template.

### [License](LICENSE)

The license covers the files in this repository *except* those under `fdroid/`.
Those APKs belong to their respective upstream projects and keep their own
licenses (Continuum: AGPL-3.0); use an F-Droid client to see per-app details.
