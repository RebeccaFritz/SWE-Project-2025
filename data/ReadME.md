#### A note on SQL and how it will interact with the server
SQL code is static, thats to say, any dynamic behavior will
be provided by our server lang using SQL commands. 

For example, the `update_wins.sql` contains SQL code to add one win
to the wins of A_Moniker22. To make this code actually useful, we would use Go to execute some similar SQL code but pass it the user we want to update. 




