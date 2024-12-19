package keys

import "embed"

// Correct as of aports revision 6d473fb38effb2389f567b29fb7eb27039b3a279

//go:embed alpine-devel@lists.alpinelinux.org-58199dcc.rsa.pub alpine-devel@lists.alpinelinux.org-616ae350.rsa.pub
var AArch64 embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-524d27bb.rsa.pub alpine-devel@lists.alpinelinux.org-616a9724.rsa.pub
var ARMhf embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-524d27bb.rsa.pub alpine-devel@lists.alpinelinux.org-616adfeb.rsa.pub
var ARMv7 embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-5243ef4b.rsa.pub alpine-devel@lists.alpinelinux.org-61666e3f.rsa.pub alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub
var X86 embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub alpine-devel@lists.alpinelinux.org-5261cecb.rsa.pub alpine-devel@lists.alpinelinux.org-6165ee59.rsa.pub
var X86_64 embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-58cbb476.rsa.pub alpine-devel@lists.alpinelinux.org-616abc23.rsa.pub
var PPC64le embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-58e4f17d.rsa.pub alpine-devel@lists.alpinelinux.org-616ac3bc.rsa.pub
var S390X embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-5e69ca50.rsa.pub
var MIPS64 embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-60ac2099.rsa.pub alpine-devel@lists.alpinelinux.org-616db30d.rsa.pub
var RISCV64 embed.FS

//go:embed alpine-devel@lists.alpinelinux.org-66ba20fe.rsa.pub
var LooongArch64 embed.FS
