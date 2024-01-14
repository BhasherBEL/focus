#! /usr/bin/env python3
from os import environ

release_tag = environ.get('RELEASE_TAG')
if not release_tag:
    print('::error ::RELEASE_TAG not set')
    exit(1)

# TODO
with open('./ci/pkgbuild/server-bin/sha256sum', 'r') as f:
    sha256sum = f.read().strip()
if not sha256sum:
    print('::error ::sha256sum not set')
    exit(1)

pkgbuild = f'''\
    # Maintainer: Brieuc Dubois <focus dot aur at bhasher dot com> 
    pkgname=focus-server-bin
    pkgver={release_tag}
    pkgrel=1
    pkgdesc="Focus is an open-source, Kanban-style project management tool, emphasizing simplicity and efficiency."
    arch=('x86_64')
    url="https://git.bhasher.com/bhasher/focus"
    license=('MIT')

    provides=("${{pkgname%-bin}}=${{pkgver}}")
    conflicts=("${{pkgname%-bin}}")

    depends=(
        'glibc'
    )

    source=(
        "focus-server::https://git.bhasher.com/Bhasher/focus/releases/download/v${{pkgver}}/focus-server"
        "LICENSE::https://git.bhasher.com/Bhasher/focus/raw/branch/master/LICENSE.md"
    )

    sha256sums=(
        '{sha256sum}'
        'SKIP'
    )

    package() {{
        install -Dm755 "${{srcdir}}/focus-server" "${{pkgdir}}/usr/bin/focus-server"
        install -Dm644 "${{srcdir}}/LICENSE" "${{pkgdir}}/usr/share/licenses/${{pkgname}}/LICENSE"
    }}
'''

with open('./ci/pkgbuild/server-bin/PKGBUILD', 'w') as f:
    f.write(pkgbuild)