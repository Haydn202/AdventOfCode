namespace DemoToDo.Models;

public class TodoItemDto
{
    public Guid Id { get; set;  }
    public string Title { get; set; }
    public bool IsComplete { get; set; }
}