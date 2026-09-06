# Setup — personal F-Droid repo (Continuum)

Base: cloned from **xarantolus/fdroid** (not forked).
`origin` now points at your own repo; the template is kept as `upstream`.

```
upstream  https://github.com/xarantolus/fdroid.git   (pull template fixes)
origin    https://github.com/<you>/fdroid.git         (your repo — set this)
```

Python tooling lives in a **local venv**: `./.venv/` → run F-Droid as `./.venv/bin/fdroid`.
No npm is needed. `metascoop` is Go and only builds inside GitHub Actions.

---

## 0. One-time local prerequisites (not Python)

`fdroid init` needs a JDK (`keytool`) and the index signer needs `apksigner`:

```bash
sudo apt-get update
sudo apt-get install -y default-jdk apksigner
keytool -help >/dev/null && apksigner --version   # both must succeed
```

(Optional, only to test metascoop locally: install Go 1.25+. CI does this for you.)

---

## 1. Wipe the template's apps, keep the plumbing  ✅ DONE

`fdroid/metadata/` and the template's demo APKs were removed; `fdroid/repo/`
is empty. `apps.yaml` now lists only **continuum** — read the multi-ABI caveat
comment in it before the first CI run.

---

## 2. Generate your repo signing key + config  ✅ DONE

`../.venv/bin/fdroid init` was run in `fdroid/`. The two `ERROR: No Android SDK`
/ `Repository already exists` lines it printed are **harmless** — the SDK is only
needed for `fdroid update` (CI does that with `apksigner` from apt). It produced:

| File | What it is | Committed? |
| --- | --- | --- |
| `fdroid/keystore.p12` | **Your repo signing key.** Signs the index only. Permanent. Valid to 2054. | No — gitignored |
| `fdroid/config.yml`   | Repo config; `keystorepass`/`keypass`/`repo_keyalias` filled in | No — gitignored |

**Repo fingerprint** (users need this to trust the repo):

```
d5c9373ac0431150dc54b59a6364e1dfb0ead9ef61591e14f95845168eb7a794
```

**Back up `fdroid/keystore.p12` offline now.** If it is ever lost, every user
must re-add a brand-new repo.

---

## 3. Edit `fdroid/config.yml`  ✅ MOSTLY DONE

Already set: `repo_name`, `repo_description`, `archive_older: 0`, and the
Continuum signing-key pin:

```yaml
allowed_apk_signing_keys:
  org.cygnusx1.continuum: 4f34bfd20cc4841510e7effb75cbfbc23a09c2ca92d04ced5b336c8f25a0d3e3
```

**One thing left** — `repo_url` has a placeholder. Replace `__GH_USER__` with
your GitHub username:

```bash
sed -i 's/__GH_USER__/YOURNAME/' fdroid/config.yml
grep '^repo_url:' fdroid/config.yml   # verify
```

The repo on GitHub must be named exactly **`fdroid`** for this URL to work.

---

## 4. Get the fingerprint + the three GitHub secrets

**Fingerprint** (users need this to trust the repo):

```bash
keytool -list -v -keystore fdroid/keystore.p12 \
  -storepass "$(grep '^keystorepass:' fdroid/config.yml | cut -d'"' -f2)" \
  2>/dev/null | grep -i 'SHA256:' | head -1
```

Take the hex, remove the colons → that is `<FINGERPRINT>`.

**Secrets** — create under *repo → Settings → Secrets and variables → Actions →
Repository secrets* (NOT environment secrets):

| Secret name | Value (command to produce it) |
| --- | --- |
| `CONFIG_YML`      | `base64 -w0 fdroid/config.yml` |
| `KEYSTORE_P12`    | `base64 -w0 fdroid/keystore.p12` |
| `GH_ACCESS_TOKEN` | A GitHub **fine-grained or classic PAT with no scopes** — used only to raise the API rate limit while metascoop polls `cygnusx-1-org/continuum` releases. No expiry, or plan to rotate it. |

---

## 5. Create and push your repo

The `repo_url` above assumes the repo is named **`fdroid`** and served from
`raw.githubusercontent.com` (no GitHub Pages needed).

```bash
git remote rename origin upstream
git remote add origin https://github.com/<you>/fdroid.git
git add -A && git commit -m "Base personal F-Droid repo on xarantolus/fdroid"
git push -u origin main
```

---

## 6. Enable and run CI

1. *Settings → Actions → General* → allow workflows to run (a pushed clone often
   needs this approved once).
2. *Actions → "Generate F-Droid repo" → Run workflow* on `main`
   (it also runs on every push and nightly at 02:45 UTC).
3. Watch the run. metascoop pulls the latest Continuum release APK, then
   `fdroid update` regenerates and signs the index and commits it back to `main`.

If the run fails on `release ... has N APK assets and none is universal`, apply
one of the two options in the `apps.yaml` comment (patch `FindAPKRelease` or
move to the manual `fdroid update` flow).

---

## 7. Add on a device

```
https://raw.githubusercontent.com/<you>/fdroid/main/fdroid/repo?fingerprint=<FINGERPRINT>
```

F-Droid → Settings → Repositories → + . Verify the fingerprint matches on first
add. Continuum then auto-updates from your repo.

---

## Keeping the template up to date

```bash
git fetch upstream
git merge upstream/main      # or cherry-pick specific fixes (e.g. metascoop)
```
