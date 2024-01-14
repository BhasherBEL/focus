#! /usr/bin/env python3
from os import environ

release_tag = environ.get('RELEASE_TAG')
if not release_tag:
    print('::error ::RELEASE_TAG not set')
    exit(1)

# TODO
with open('./ci/pkgbuild/desktop-bin/sha256sum', 'r') as f:
    sha256sum = f.read().strip()
if not sha256sum:
    print('::error ::sha256sum not set')
    exit(1)

pkgbuild = f'''\
    # Maintainer: Brieuc Dubois <focus dot aur at bhasher dot com> 
    pkgname=focus-desktop-bin
    pkgver={release_tag}
    pkgrel=1
    pkgdesc="Focus is an open-source, Kanban-style project management tool, emphasizing simplicity and efficiency."
    arch=('x86_64')
    url="https://git.bhasher.com/bhasher/focus"
    license=('MIT')

    provides=("${{pkgname%-bin}}=${{pkgver}}")
    conflicts=("${{pkgname%-bin}}")

    depends=(
        'webkit2gtk'
        'gtk3'
        'cairo'
        'glib2'
        'hicolor-icon-theme'
        'gdk-pixbuf2'
        'libsoup'
        'gcc-libs'
        'glibc'
        'pango'
    )

    source=(
        "focus-desktop_${{pkgver}}_amd64.deb::https://git.bhasher.com/Bhasher/focus/releases/download/v${{pkgver}}/focus-desktop_${{pkgver}}_amd64.deb"
        "LICENSE::https://git.bhasher.com/Bhasher/focus/raw/branch/master/LICENSE.md"
    )

    sha256sums=(
        '{sha256sum}'
        'SKIP'
    )

    build() {{
        bsdtar -xf "${{srcdir}}/data.tar.gz"
    }}

    package() {{
        install -Dm755 "${{srcdir}}/usr/bin/focus-desktop" "${{pkgdir}}/usr/bin/focus-desktop"
        install -Dm644 "${{srcdir}}/usr/share/applications/focus-desktop.desktop" "${{pkgdir}}/usr/share/applications/focus-desktop.desktop"
        for _icons in 32x32 128x128 256x256@2;do
            install -Dm644 "${{srcdir}}/usr/share/icons/hicolor/${{_icons}}/apps/focus-desktop.png" \
                -t "${{pkgdir}}/usr/share/icons/hicolor/${{_icons//@2/}}/apps/focus-desktop.png"
        done
        install -Dm644 "${{srcdir}}/LICENSE" "${{pkgdir}}/usr/share/licenses/${{pkgname}}/LICENSE"
    }}
'''

with open('./ci/pkgbuild/desktop-bin/PKGBUILD', 'w') as f:
    f.write(pkgbuild)