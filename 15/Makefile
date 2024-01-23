BOCHS	= ../bochs/
NASM 	= nasm -f bin
MV 		= mv
MAKE	= make
DEL     = rm
DD		= dd
CAT		= cat

default:
	$(NASM) ../13/c13_mbr.asm -l c13_mbr.lst -o c13_mbr.bin
	$(NASM) c15_core.asm -l c15_core.lst -o c15_core.bin
	$(NASM) c15.asm -l c15.lst -o c15.bin

build: 
	$(MAKE)
	$(DD) if=/dev/zero of=$(BOCHS)fd.img bs=1024 count=10080
	$(DD) if=c13_mbr.bin of=$(BOCHS)fd.img bs=512 count=1 conv=notrunc
	$(DD) if=c15_core.bin of=$(BOCHS)fd.img bs=512 seek=1 conv=notrunc
	$(DD) if=c15.bin of=$(BOCHS)fd.img bs=512 seek=50 conv=notrunc
	$(DD) if=../13/diskdata.txt of=$(BOCHS)fd.img bs=512 seek=100 conv=notrunc

run:	
	$(MAKE) build
	$(MAKE) -C $(BOCHS)

clean:
	-$(DEL) *.bin
	-$(DEL) *.lst
