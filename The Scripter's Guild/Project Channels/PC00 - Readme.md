
Purpose:
*Provides infrastructure for managing collaborative projects in Discord. Supports having multiple 'project groups' spread between a channel per group.*

Dependencies:
 - *Dyno: Makes use of Dyno's Forms module. The form output is the initial input for this system.
 - A bot that can handle autothreading and will place threads on bot messages. I use Needle.
 - A channel for the form output.
 - A channel for each project group (unless you want them all in the same channel for some reason).
 - Autothreading configured on each project channel.

 *Technically, you could use another input, but you would need to rewrite **PC02** to take that input and produce the same output.

**PC01** `-projectCount`

Use this command to see and manipulate project count db entries.

Syntax: `-projectCount <subcommand> <target> [value]`

Subcommands: `get`, `set`, `help`

`get`: prints project counts of all or target project group(s)

`set`: sets project count of target project group and prints info about the change

`help`: prints help message

Target: `all`, `<projectGroup>`

`all`: only works with `get`, causes it to print all groups' counts
`<projectGroup>`: must match one of the project groups defined in **PC01**

	Groups and their reference syntax for `<projectGroup>`:

	Infinite	->	`infinite` 
	Legacy Halo	->	`legacyHalo`
	Unreal		->	`unreal`
	Unity		->	`unity`
	Core		->	`core`
	Fortnite	->	`fortnite`

Value: an integer, the number you want to set the relevant projecy count db value to, is not used when using `get` or `help` subcommands

**PC02** `Added Reaction in #📥｜forms-output (982661501766754355)`

When the form output embed shows up in #📥｜forms-output, Dyno automatically adds 👍 and 👎 reactions to the message.

Adding an additional of these emoji reactions to the embed triggers this command, as it watches for the count for that emoji to reach 2. 

2x 👍: Approved. Existing emoji are cleared and replaced with ✅, approval code processes.

2x 👎: Denied. Existing emoji are cleared and replaced with ❎.

2x ✅ or ❎: Reset: Existing emoji are cleared and replaced with 👍 and 👎.

Approval:

1. Embed is copied in place to turn Dyno's custom emoji into text that can be easily regex'd.

2. Data is parsed from copy to determine where to send the new project embed.

3. Copy is deleted, new project embed is sent. 

4. Needle automatically creates a thread attached to that embed message and sends a message inside that thread to explain how to use project channel relevant commands.

Note: 
*Unfotunately, the name of the thread created is not able to be automatically changed. This is resolved when the user first runs `-project` and the thread's name is updated to match the project name. This would handle that if it where able.*

**PC03** `-project`

Use this command to update/edit your Project Embed. You have to properly target the embed and it has to be done from the relevant thread for the commands to work. You must be a lead on the relvant project or a Guild Officer to use this command.

Note: *This command must be set to *only* work in the channels that house project threads.*

Syntax:

`-project <[embedMessageID]|[help]> <subCommand> <value>`

Performs given subcommand on targeted embed with given value.

`<[embedMessageID]|[help]>`: Embed Message ID. Copy this from the footer of the Project Embed. Also accepts `help` if only argument given.

`<subCommand>`: Subcommand. Each valueAccepted values: `version`, `name`, `description`, `links`, `addLead`, `removeLead`, `addContributor`, and `removeContributor`.

`<value>`: Relevant input per Subcommand. `description` and `links` subcommands accept normal message formatting and as a line break, but you can just add line breaks in your command message. Enclose in double quotes "like this" if you use any spaces in this argument.



