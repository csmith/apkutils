package keys

import (
	"embed"
	"github.com/csmith/apkutils"
)

// Correct as of aports revision 6d473fb38effb2389f567b29fb7eb27039b3a279

//go:embed alpine-devel@lists.alpinelinux.org-58199dcc.rsa.pub alpine-devel@lists.alpinelinux.org-616ae350.rsa.pub
var aarch64 embed.FS
var Aarch64 = apkutils.NewFileSystemKeyProvider(aarch64)

//go:embed alpine-devel@lists.alpinelinux.org-524d27bb.rsa.pub alpine-devel@lists.alpinelinux.org-616a9724.rsa.pub
var armhf embed.FS
var ARMhf = apkutils.NewFileSystemKeyProvider(armhf)

//go:embed alpine-devel@lists.alpinelinux.org-524d27bb.rsa.pub alpine-devel@lists.alpinelinux.org-616adfeb.rsa.pub
var armv7 embed.FS
var ARMV7 = apkutils.NewFileSystemKeyProvider(armv7)

//go:embed alpine-devel@lists.alpinelinux.org-5243ef4b.rsa.pub alpine-devel@lists.alpinelinux.org-61666e3f.rsa.pub alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub
var x86 embed.FS
var X86 = apkutils.NewFileSystemKeyProvider(x86)

//go:embed alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub alpine-devel@lists.alpinelinux.org-5261cecb.rsa.pub alpine-devel@lists.alpinelinux.org-6165ee59.rsa.pub
var x86_64 embed.FS
var X86_64 = apkutils.NewFileSystemKeyProvider(x86_64)

//go:embed alpine-devel@lists.alpinelinux.org-58cbb476.rsa.pub alpine-devel@lists.alpinelinux.org-616abc23.rsa.pub
var ppc64le embed.FS
var PPC64le = apkutils.NewFileSystemKeyProvider(ppc64le)

//go:embed alpine-devel@lists.alpinelinux.org-58e4f17d.rsa.pub alpine-devel@lists.alpinelinux.org-616ac3bc.rsa.pub
var s390x embed.FS
var S390X = apkutils.NewFileSystemKeyProvider(s390x)

//go:embed alpine-devel@lists.alpinelinux.org-5e69ca50.rsa.pub
var mips64 embed.FS
var MIPS64 = apkutils.NewFileSystemKeyProvider(mips64)

//go:embed alpine-devel@lists.alpinelinux.org-60ac2099.rsa.pub alpine-devel@lists.alpinelinux.org-616db30d.rsa.pub
var riscv64 embed.FS
var RISCV64 = apkutils.NewFileSystemKeyProvider(riscv64)

//go:embed alpine-devel@lists.alpinelinux.org-66ba20fe.rsa.pub
var looongarch64 embed.FS
var LooongArch64 = apkutils.NewFileSystemKeyProvider(looongarch64)

//go:embed *.rsa.pub
var all embed.FS
var All = apkutils.NewFileSystemKeyProvider(all)
