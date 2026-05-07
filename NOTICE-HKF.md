# Notice — HKF Fork of NewAPI

This is a fork of [QuantumNous/new-api](https://github.com/QuantumNous/new-api) (originally by [Calcium-Ion](https://github.com/Calcium-Ion/new-api), now maintained by QuantumNous).

## What we changed

Visual adaptation to HKF (Houkongfan AI Gateway) warm color palette:

- `web/classic/src/index.css` — `.blur-ball-indigo` color #6366f1 → #C2410C; `.blur-ball-teal` color #14b8a6 → #F59E0B
- `web/classic/src/components/auth/{LoginForm,RegisterForm,PasswordResetForm,PasswordResetConfirm}.jsx` — outer wrapper `bg-gray-100` → transparent, letting the body background show through

## What we did NOT change

Per upstream `CLAUDE.md` Rule 5, all references to **NewAPI** and **QuantumNous** are preserved unchanged in:

- README files, license headers, copyright notices
- Footer attribution ("设计与开发由 New API")
- HTML titles, meta tags
- Go module paths, package names
- Docker image names in upstream Dockerfile / CI

## License

This fork inherits **AGPL-3.0** from upstream. See `LICENSE` for the full text.

Per AGPL §13 (network use), the source code of this modified version is published at:
**https://github.com/beikeint/new-api-hkf**

For commercial licensing of upstream NewAPI, contact `support@quantumnous.com` (per upstream README).

## Contact (this fork)

- Operator: 江阴市后空翻网络科技有限公司
- Email: lilong1800@gmail.com
