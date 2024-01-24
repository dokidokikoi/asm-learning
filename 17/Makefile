BOCHS	= ../bochs/
NASM 	= nasm -f bin
MV 		= mv
MAKE	= make
DEL     = rm
DD		= dd
CAT		= cat

default:
	$(NASM) c17_mbr.asm -l c17_mbr.lst -o c17_mbr.bin
	$(NASM) c17_core.asm -l c17_core.lst -o c17_core.bin
	$(NASM) c17-1.asm -l c17-1.lst -o c17-1.bin
	$(NASM) c17-2.asm -l c17-2.lst -o c17-2.bin

build: 
	$(MAKE)
	$(DD) if=/dev/zero of=$(BOCHS)fd.img bs=1024 count=10080
	$(DD) if=c17_mbr.bin of=$(BOCHS)fd.img bs=512 count=1 conv=notrunc
	$(DD) if=c17_core.bin of=$(BOCHS)fd.img bs=512 seek=1 conv=notrunc
	$(DD) if=c17-1.bin of=$(BOCHS)fd.img bs=512 seek=50 conv=notrunc
	$(DD) if=c17-2.bin of=$(BOCHS)fd.img bs=512 seek=100 conv=notrunc

run:	
	$(MAKE) build
	$(MAKE) -C $(BOCHS)

clean:
	-$(DEL) *.bin
	-$(DEL) *.lst
