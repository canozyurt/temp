FROM    mcr.microsoft.com/dotnet/core/sdk:3.1
RUN mkdir /app

WORKDIR /app
COPY    . .

RUN     dotnet publish YksQuestions.WebAPI.csproj -o /publish 

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-bionic

ENV     LANG tr_TR.utf8
ENV     TZ 'Europe/Istanbul'

RUN apt update && \
    apt-get install -y tzdata locales && \
            localedef -i tr_TR -c -f UTF-8 -A /usr/share/locale/locale.alias $LANG && \
            echo $TZ > /etc/timezone && \
            rm /etc/localtime && \
            ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
            dpkg-reconfigure -f noninteractive tzdata && \
            rm -rf /var/lib/apt/lists/* && apt-get clean
      
ENV ASPNETCORE_URLS=http://*:5000 \
    ASPNETCORE_ENVIRONMENT=Development

WORKDIR /app

COPY --from=0 ./publish /app

ENTRYPOINT ["dotnet", "YksQuestions.WebAPI.dll"]