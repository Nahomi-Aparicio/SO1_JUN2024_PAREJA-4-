#include <linux/module.h>
#define INCLUDE_VERMAGIC
#include <linux/build-salt.h>
#include <linux/elfnote-lto.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

BUILD_SALT;
BUILD_LTO_INFO;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(".gnu.linkonce.this_module") = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};

#ifdef CONFIG_RETPOLINE
MODULE_INFO(retpoline, "Y");
#endif

static const struct modversion_info ____versions[]
__used __section("__versions") = {
	{ 0x602c1205, "module_layout" },
	{ 0xacb46524, "seq_read" },
	{ 0x363f205f, "remove_proc_entry" },
	{ 0x4334b731, "proc_create" },
	{ 0xb32ec28b, "from_kuid" },
	{ 0x6fbcc169, "init_user_ns" },
	{ 0x8d286714, "seq_printf" },
	{ 0x37befc70, "jiffies_to_msecs" },
	{ 0x825fce9b, "init_task" },
	{ 0x944375db, "_totalram_pages" },
	{ 0x92997ed8, "_printk" },
	{ 0xd0da656b, "__stack_chk_fail" },
	{ 0x4f418e78, "filp_close" },
	{ 0xbcab6ee6, "sscanf" },
	{ 0xc69def87, "kernel_read" },
	{ 0x9987d206, "filp_open" },
	{ 0x5b8239ca, "__x86_return_thunk" },
	{ 0x2dcf0dc0, "single_open" },
	{ 0xbdfb6dbb, "__fentry__" },
};

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "C26ABA1443C6CB141DCD463");
