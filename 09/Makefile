BOCHS	= ../bochs/
NASM 	= nasm -f bin
MV 		= mv
MAKE	= make
DEL     = rm
DD		= dd
CAT		= cat

default:
	$(NASM) ../08/c08_mbr.asm -o c08_mbr.bin
	$(NASM) c09_1.asm -l c09_1.lst -o c09_1.bin
	$(NASM) c09_2.asm -l c09_2.lst -o c09_2.bin

build: 
	$(MAKE)
	$(DD) if=/dev/zero of=$(BOCHS)fd.img bs=1024 count=10080
	$(DD) if=c08_mbr.bin of=$(BOCHS)fd.img bs=512 count=1 conv=notrunc

run1:	
	$(MAKE) build
	$(DD) if=c09_1.bin of=$(BOCHS)fd.img bs=512 seek=100 conv=notrunc
	$(MAKE) -C $(BOCHS)

run2:	
	$(MAKE) build
	$(DD) if=c09_2.bin of=$(BOCHS)fd.img bs=512 seek=100 conv=notrunc
	$(MAKE) -C $(BOCHS)

clean:
	-$(DEL) *.bin
	-$(DEL) *.lst