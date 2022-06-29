This utility will add everyone on your server to threads that the "-share" command is run in. 
The reason you would want to do this is for threads that may be unarchived as they will only show in the sidebar for 
people who already interacted with them or were tagged in them. This performs a silent ping to add users in 10 roles
of up to 100 users each to a thread.

Note: This only works for servers with less that 1k members because of limitations with Threads.

Create a "Sort" Role and have it applied to all people who join.
Create 10 roles: 100, 200...900, 1000. 
Put every existing member in **1** of those roles, filling up the lowest numbered ones first, in order.

Setup:
Use the setup command to set the database value to your current member count. 
You can use this stored member count db entry for other commands, now, too (even if you have more than 1000 people).

Reaction&Regex: 

Set up a command each with a Reaction and Regex trigger, respectively. Same code.
Regex: .*
Reaction: Added and removed, require "Sort" Role.
Change Role ID in removeRoleID line to match your "Sort" Role.
Change each Role ID in the if block to the ID of the Role who's name matches the value in the relevant if/if else. 

Share:
Set it up with  so that users can run the "-share" command in channels where they can make threads (and you want them 
to be able to add people to them like this).
Change each Role ID in the messageEdits to the ID of the 10 count roles, respectively.
