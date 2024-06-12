#include <linux/module.h> // THIS_MODULE, MODULE_VERSION, ...
#include <linux/init.h>   // module_{init,exit}
#include <linux/proc_fs.h>
#include <linux/sched/signal.h> // for_each_process()
#include <linux/seq_file.h>
#include <linux/fs.h>
#include <linux/sched.h>
#include <linux/mm.h> // get_mm_rss()
#include <linux/cred.h> // struct cred
#include <linux/uaccess.h> // for kernel_read()

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Nahomi Aparicio -Richard Marroquin");
MODULE_DESCRIPTION("Modulo de CPU para el Lab de Sopes 1");


struct task_struct *cpu;       // sched.h para tareas/procesos
struct list_head *lstProcess;
unsigned long rss;

static int calcularPorcentajeCpu(void)
{
    struct file *file_proc; // Puntero a estructura de archivo para manejar el archivo /proc/stat
    char lectura[256]; // Array de caracteres para almacenar la lectura del archivo /proc/stat

    int usuario, nice, system, idle, iowait, irq, softirq, steal, guest, guest_nice;
    int total;
    int porcentaje;

    // Abre el archivo /proc/stat en modo lectura
    file_proc = filp_open("/proc/stat", O_RDONLY, 0);
    if (IS_ERR(file_proc)) // IS_ERR se usa para verificar si hay un error al abrir el archivo
    {
        printk(KERN_ALERT "Error al abrir el file_proc");
        return -1;
    }

    memset(lectura, 0, 256);
    kernel_read(file_proc, lectura, sizeof(lectura), &file_proc->f_pos);

    // Extracción de los valores de uso de CPU del archivo
    sscanf(lectura, "cpu %d %d %d %d %d %d %d %d %d %d", &usuario, &nice, &system, &idle, &iowait, &irq, &softirq, &steal, &guest, &guest_nice);
    total = usuario + nice + system + idle + iowait + irq + softirq + steal + guest + guest_nice;

    // Calcula el porcentaje utilizando matemáticas enteras
    if (total > 0) {
        porcentaje = (total - idle) * 1000 / total; // Multiplica por 1000 para mayor precisión y evita punto flotante
    } else {
        porcentaje = 0;
    }

    // Cierra el archivo
    filp_close(file_proc, NULL);

    return porcentaje;
}

static int escribir_a_proc(struct seq_file *archivo, void *v)
{
    struct task_struct *child; // Declarar child antes de usarlo
    int porcentajecpu = calcularPorcentajeCpu();

    // si retorna un -1 hay un error en el archivo 
    if (porcentajecpu == -1)
    {
        seq_printf(archivo, "Error al leer el archivo");
        return 0;
    }

    // Escribo el porcentaje de CPU que se esta usando 
    seq_printf(archivo, "{\"cpu_porcentaje\":%d.%d ,\n", porcentajecpu / 10, porcentajecpu % 10); // Formatear el porcentaje correctamente

   seq_printf(archivo, "\"Procesos_existentes\":[\n");
int b = 0;

for_each_process(cpu) {
    if (cpu->mm)
    {
        rss = get_mm_rss(cpu->mm) << PAGE_SHIFT;
    }
    else
    {
        rss = 0;
    }
    if (b == 0)
    {
        seq_printf(archivo, "{");
        b = 1;
    }
    else
    {
        seq_printf(archivo, ",{");
    }
    seq_printf(archivo, "\"pid\":%d,\n", cpu->pid);
    seq_printf(archivo, "\"name\":\"%s\",\n", cpu->comm);
    seq_printf(archivo, "\"state\":%u,\n", cpu->__state);
    seq_printf(archivo, "\"rss\":%lu,\n", rss);
    seq_printf(archivo, "\"uid\":%u,\n", from_kuid(&init_user_ns, cpu->cred->uid));

    seq_printf(archivo, "\"children\":[\n");
    int a = 0;
    struct list_head *list;
    struct task_struct *task_child;
    list_for_each(list, &(cpu->children))
    {
        task_child = list_entry(list, struct task_struct, sibling);
        if (a != 0)
        {
            seq_printf(archivo, ",{");
        }
        else
        {
            seq_printf(archivo, "{");
            a = 1;
        }
        seq_printf(archivo, "\"pid\":%d,\n", task_child->pid);
        seq_printf(archivo, "\"name\":\"%s\",\n", task_child->comm);
        seq_printf(archivo, "\"state\":%u,\n", task_child->__state);
        seq_printf(archivo, "\"pidPadre\":%d,\n", cpu->pid);
        if (task_child->mm)
        {
            rss = get_mm_rss(task_child->mm) << PAGE_SHIFT;
        }
        else
        {
            rss = 0;
        }
        seq_printf(archivo, "\"rss\":%lu,\n", rss);
        seq_printf(archivo, "\"uid\":%u\n", from_kuid(&init_user_ns, task_child->cred->uid));
        seq_printf(archivo, "}\n");
    }
    seq_printf(archivo, "]\n");
    seq_printf(archivo, "}\n");
}
b = 0;
seq_printf(archivo, "]\n");
seq_printf(archivo, "}\n");

    return 0;
}

static int abrir_aproc(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_a_proc, NULL);
}

static struct proc_ops archivo_operaciones = {
    .proc_open = abrir_aproc,
    .proc_read = seq_read
};

static int __init modulo_init(void)
{
    proc_create("cpu_so1_jun2024", 0, NULL, &archivo_operaciones);
    printk(KERN_INFO "Nahomi Aparicio -Richard Marroquin\nn");
    return 0;
}

static void __exit modulo_cleanup(void)
{
    remove_proc_entry("cpu_so1_jun2024", NULL);
    printk(KERN_INFO "Esuela de vacaiones 2024\n");
}

module_init(modulo_init);
module_exit(modulo_cleanup);

