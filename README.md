# Parsing-json-page-using-Go
Use 3rd-party JSON API: https://date.nager.at/PublicHoliday/Country/UA
Write a console application that prints if it’s a holiday today (and the name of it).
If today isn’t a holiday, the application should print the next closest holiday. 

Additionally, if the holiday is adjacent to a weekend (so that amount of non-working days is extended), 
the application should print this information. I.e. the next holiday is May 1, Friday, and it’s adjacent
to Saturday (May 2) and Sunday (May 3), so the application should print something like: 
“The next holiday is International Workers' Day, May 1, and the weekend will last 3 days: May 1 - May 3”.
