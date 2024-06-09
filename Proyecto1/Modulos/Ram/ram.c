#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sysinfo.h> // Para obtener información de la RAM
#include <linux/seq_file.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM, Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Nahomi Aparicio - Richard Marroquin");

static struct sysinfo inf;

// Función para obtener y formatear la información de la memoria
static int mostrar_info_ram(struct seq_file *file_proc, void *v)
{
    unsigned long total, usado;
    unsigned long porcentaje_usado;

    si_meminfo(&inf); // Obtener información de la memoria

    total = inf.totalram * inf.mem_unit;
   
    usado = inf.freeram * inf.mem_unit + inf.bufferram * inf.mem_unit + inf.sharedram * inf.mem_unit;
    porcentaje_usado = (usado * 100) / total;

    // Escribir la información en formato JSON en el archivo proc
    seq_printf(file_proc, "{\"totalRam\":%lu, \"memoriaEnUso\":%lu, \"porcentaje\":%lu}", 
               total, usado, porcentaje_usado);
    return 0;
}

// Función para abrir el archivo proc
static int abrir_info_ram(struct inode *inode, struct file *file)
{
    return single_open(file, mostrar_info_ram, NULL);
}

// Operaciones del archivo proc para kernels superiores a la versión 5.6
static const struct proc_ops proc_ops_info_ram = {
    .proc_open = abrir_info_ram,
    .proc_read = seq_read,
    .proc_lseek = seq_lseek,
    .proc_release = single_release
};

// Función de inicialización del módulo
static int __init modulo_init(void)
{
    proc_create("ram_so1_jun2024", 0, NULL, &proc_ops_info_ram);
    printk(KERN_INFO "Modulo RAM montado\n");
    return 0;
}

// Función de limpieza del módulo
static void __exit modulo_cleanup(void)
{
    remove_proc_entry("ram_so1_jun2024", NULL);
    printk(KERN_INFO "Modulo RAM eliminado\n");
}

module_init(modulo_init);
module_exit(modulo_cleanup);

