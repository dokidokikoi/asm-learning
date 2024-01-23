BOCHS	= ../bochs/
NASM 	= nasm -f bin
MV 		= mv
MAKE	= make
DEL     = rm
DD		= dd
CAT		= cat

default:
	$(NASM) ../13/c13_mbr.asm -l c13_mbr.lst -o c13_mbr.bin
	$(NASM) c16_core.asm -l c16_core.lst -o c16_core.bin
	$(NASM) c16.asm -l c16.lst -o c16.bin

build: 
	$(MAKE)
	$(DD) if=/dev/zero of=$(BOCHS)fd.img bs=1024 count=10080
	$(DD) if=c13_mbr.bin of=$(BOCHS)fd.img bs=512 count=1 conv=notrunc
	$(DD) if=c16_core.bin of=$(BOCHS)fd.img bs=512 seek=1 conv=notrunc
	$(DD) if=c16.bin of=$(BOCHS)fd.img bs=512 seek=50 conv=notrunc

run:	
	$(MAKE) build
	$(MAKE) -C $(BOCHS)

clean:
	-$(DEL) *.bin
	-$(DEL) *.lst
