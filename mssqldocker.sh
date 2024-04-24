docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=1VerySecurePassword!" -e "MSSQL_PID=Developer" -p 1433:1433 --rm  --name mssql --hostname mssql -d mcr.microsoft.com/mssql/server:2022-latest
