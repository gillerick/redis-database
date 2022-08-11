Implement an in-memory database that has the following functions. We’ll be looking for your code to
meet our functionality & performance requirements, be clear & easy to understand and be resilient to
edge cases. Use libraries at will (but not database-like ones or actual databases). Use Google/Stack
Overflow/online references at will as well.

The database should be a command line program that reads values from STDIN line by line and executes
the functions as they happen. Please also include a README explaining how to run your program.

The name and value will be strings with no spaces in them.

SET [name] [value]
GET [name]
DELETE [name]
COUNT [value]
END
BEGIN
ROLLBACK
COMMIT