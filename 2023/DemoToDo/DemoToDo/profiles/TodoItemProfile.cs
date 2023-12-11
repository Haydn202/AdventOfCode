using AutoMapper;
using DemoToDo.Entities;
using DemoToDo.Models;

namespace DemoToDo.profiles;

public class TodoItemProfile : Profile
{
    public TodoItemProfile()
    {
        CreateMap<TodoItem, TodoItemDto>();
        CreateMap<TodoItemCreationDto, TodoItem>();
    }
}