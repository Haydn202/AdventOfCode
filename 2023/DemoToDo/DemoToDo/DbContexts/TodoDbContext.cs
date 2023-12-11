using DemoToDo.Entities;
using Microsoft.EntityFrameworkCore;

namespace DemoToDo.DbContexts;

public class TodoDbContext: DbContext
{
    public TodoDbContext(DbContextOptions<TodoDbContext> options) : base(options) { }

    public DbSet<TodoItem> TodoItems { get; set; }
}