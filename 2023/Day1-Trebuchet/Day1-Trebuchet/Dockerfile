FROM mcr.microsoft.com/dotnet/runtime:8.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["Day1-Trebuchet/Day1-Trebuchet.csproj", "Day1-Trebuchet/"]
RUN dotnet restore "Day1-Trebuchet/Day1-Trebuchet.csproj"
COPY . .
WORKDIR "/src/Day1-Trebuchet"
RUN dotnet build "Day1-Trebuchet.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Day1-Trebuchet.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Day1-Trebuchet.dll"]
