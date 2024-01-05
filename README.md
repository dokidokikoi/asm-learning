# asm-learning

x86汇编语言从实模式到保护模式 (李忠) 的实验代码及环境。

不需要像书上一样使用 virtualbox 和 bochs，仅使用 bochs 即可。不用在意软盘和硬盘的格式，直接在章节目录下运行 `make run` 即可启动 bochs，然后就可以在 bochs 中 debug 和查看运行结果了。 

**当然，你需要提前安装好 bochs、nasm 和 make 工具，并配置好环境变量**

Windows 环境下需要修改各目录下的 Makefile。