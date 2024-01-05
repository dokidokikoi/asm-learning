BOCHS	= ../bochs/
NASM 	= nasm -f bin
MV 		= mv
MAKE	= make
DEL     = rm
DD		= dd
CAT		= cat

default:
	$(NASM) c12_mbr.asm -l c12_mbr.lst -o c12_mbr.bin

build: 
	$(MAKE)
	$(DD) if=/dev/zero of=$(BOCHS)fd.img bs=1024 count=10080
	$(DD) if=c12_mbr.bin of=$(BOCHS)fd.img bs=512 count=1 conv=notrunc

run:	
	$(MAKE) build
	$(MAKE) -C $(BOCHS)

clean:
	-$(DEL) *.bin
	-$(DEL) *.lst