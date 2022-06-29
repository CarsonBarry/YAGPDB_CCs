{{ if not   (or 
				(eq (len .CmdArgs) 1) 
				(eq (len .CmdArgs) 4)
				(eq (len .CmdArgs) 5) 
			) }}

	{{ sendMessage nil "Syntax is `-embedUpdate <embedChannelID> <embedMessageID> <subCommand> <value>`, `-embedUpdate <embedChannelID> <embedMessageID> <fieldChange> <fieldIndex> <value>`, or `-embedUpdate help`" }}

{{ else if (eq (len .CmdArgs) 1) }}

	{{ $args := parseArgs 1 "Syntax is -embedUpdate <help>"
		(carg "string" "Accepts 'help' if only argument given.")
	}}

	{{ if (eq ($args.Get 0) "help") }}

		{{ $helpMessage := cembed
			"color" 000000
			"author" (sdict	"name" ("TSG Admin Team") "url" ("https://discord.gg/ZpE3tcpZy9") "icon_url" ("https://i.imgur.com/505zjJl.png"))
			"thumbnail" (sdict "url" ("https://i.imgur.com/a7NAJZU.png") )
			"title" ("Project Command Help")
			"description" (print
				"*Use this command to update/edit an Embed that was created by <@204255221017214977>. You have to properly target the embed. You must be a Guild Officer to use this command.*\n\n"
				"Syntax:\n\n"
				"**1 Argument Given**\n"
				"`-embedUpdate <help>`\n*displays this message*\n\n"
				"**4 Arguments Given**\n"
				"`-embedUpdate <embedChannelID> <embedMessageID> <subCommand> <value>`\nPerforms given subcommand on targeted embed with given value.\n\n"
				"**embedChannelID** - *Channel ID of the channel where the message is. Right click the channel name in the sidebar and select 'Copy ID' (requires dev mode to be on). Also accepts 'help' if only argument given.*\n\n"
				"**embedMessageID** - *Message ID of the message to which the embed is attached. Copy this from the footer of the embed or right click the message and select 'Copy ID' (requires dev mode to be on).*\n\n"
				"**subCommand** - *Subcommand you are trying to execute. Accepted values: author, thumbnail, title, description, fieldChange, footer*\n\n"
				"**value** - *Relevant input per Subcommand. 'description' and 'links' subcommands accept normal message formatting and \\n as a line break, but you can just add line breaks in your command message.*\n\n"
				"**5 Arguments Given**\n"
				"`-embedUpdate <embedChannelID> <embedMessageID> <fieldChange> <fieldIndex> <value>`\nPerforms given subcommand on targeted embed with given value.\n\n"
				"**fieldChange** - *The part of the field you are trying to change. Accepted values: `fieldName`, `fieldValue`, `fieldInline`.*"
				"**fieldIndex** - *The index of the field you are trying to update. Starts at 0 and runs top to bottom (left to right when inline fields are involved).*"
			)
		}}

		{{ sendMessage nil $helpMessage}}

	{{ else }}

		{{ sendMessage nil "Use `-embedUpdate help` to see the full syntax and an explanation of this command."}}

	{{ end }}

{{ else if (or 
	(eq (len .CmdArgs) 4) 
	(eq (len .CmdArgs) 5)
)}}

	{{ $args := parseArgs 4 "Syntax is -embedUpdate <embedChannelID> <embedMessageID> <subCommand> [fieldIndex] <value>"
		(carg "string" "Channel ID of the channel where the message is. Copy this from the footer of the Project Embed.")	
		(carg "string" "Message ID of the message to which the embed is attached. Copy this from the footer of the Project Embed.")
		(carg "string" "Subcommand. Accepted values: `author`, `thumbnail`, `title`, `description`, `field`, `footer`.\nUsing `field` induces the requirement to provide a `fieldIndex` prior to the `value`.")
		(carg "string" "Relevant input per Subcommand. 'description' and 'links' subcommands accept normal message formatting and \\n as a line break, but you can just add line breaks in your command message. Enclose this argument in double quotes for a more consistent experience.")
	}}

	{{/* Package Args into Variables */}}
	{{ $embedChannelID			:= ($args.Get 0) }}
	{{ $embedMessageID			:= ($args.Get 1) }}
	{{ $subCommand 				:= ($args.Get 2) }}
	{{ $value		   		:= ($args.Get 3) }}

	{{/* Get embed */}}
	{{ $embedOld			   	:= (index (getMessage $embedChannelID $embedMessageID).Embeds 0) }}


	{{/* Unpack Embed into variables */}}
	{{ $authorName				:= str $embedOld.Author.Name }}
	{{ $authorURL				:= str $embedOld.Author.URL }}
	{{ $authorIcon				:= str $embedOld.Author.IconURL }}
	{{ $embedTitle 				:= str $embedOld.Title }}
	{{ $embedDescription		:= str $embedOld.Description }}
	{{ $embedThumbnail 			:= str $embedOld.Thumbnail }}
	{{ $embedFooter 			:= $embedOld.Footer }}
    {{ $embedFooterText         := $embedFooter.Text }}
	
	{{/* Initialize/define variables */}}
	{{ $numFields := len $embedOld.Fields }}
	{{ $fields := cslice "temp" }}

	{{/* If: there is at least 1 Field */}}
	{{ if ( gt $numFields 0 ) }}

		{{/* Store that field's value as index 0 of the \$fields slice */}}
		{{ $fields = cslice (index $embedOld.Fields 0) }}

		{{/* For each: Field after the first */}}
		{{ $i := 1 }} {{ while ( lt $i $numFields ) }}
			{{/* Append next field to $fields slice */}}
			{{ $fields = $fields.Append (index $embedOld.Fields $i) }}
			{{/* Iterate Loop */}}
			{{ $i = ( add $i 1 ) }}
		{{ end }}

	{{ else }}
	{{ end }}

	{{/* The If-Switch */}}
	{{ if ( eq $subCommand "author") }}
		{{ $authorName = $value }}
	{{ else if ( eq $subCommand "thumbnail") }}
		{{ $embedThumbnail = $value }}
	{{ else if ( eq $subCommand "title") }}  
		{{ $embedTitle = $value }}
	{{ else if ( eq $subCommand "description") }}  
		{{ $embedDescription = $value }}
	{{ else if ( eq $subCommand "footer") }}  
		{{ $embedFooter = (sdict "text" $embedFooterText) }}
	{{ else if ( or ( eq $subCommand "fieldName" ) ( eq $subCommand "fieldValue" ) ( eq $subCommand "fieldInline" ) ) }}
		{{ if (eq (len .CmdArgs) 5) }}
			{{ $args := parseArgs 5 "Syntax is -embedUpdate <embedChannelID> <embedMessageID> <subCommand> <fieldIndex> <value>"
				(carg "string" "Channel ID of the channel where the message is. Copy this from the footer of the Project Embed.")	
				(carg "string" "Message ID of the message to which the embed is attached. Copy this from the footer of the Project Embed.")
				(carg "string" "When using 5 arguments, this *must* be `field`.")
				(carg "int" "The index of the field that you are trying to update.  Starts at 0 and ascends top to bottom (left to right when inline fields are involved).")
				(carg "string" "Relevant input per Subcommand. 'description' and 'links' subcommands accept normal message formatting and \\n as a line break, but you can just add line breaks in your command message. Enclose this argument in double quotes for a more consistent experience.")
			}}  
			{{/* Package Args into Variables */}}
			{{ $fieldIndex	 	:= ($args.Get 3) }}
			{{ $value			= ($args.Get 4) }}

			{{/* Unpack relevant field into variables */}}
			{{ $fieldName	 	:= ( index $fields $fieldIndex ).Name }}
			{{ $fieldValue	 	:= ( index $fields $fieldIndex ).Value }}
			{{ $fieldInline	    := ( index $fields $fieldIndex ).Inline }}

			{{/* Update relevant field...field with <value> */}}
			{{ if ( eq $subCommand "fieldName" ) }}
				{{ $fieldName	= $value }}
			{{ else if ( eq $subCommand "fieldValue" ) }}
				{{ $fieldValue   = $value }}
			{{ else }} {{/* If: fieldInline */}}
				{{ if ( eq $value "true" ) }}
					{{ $fieldInline  = true }}
                {{ else if ( eq $value "false" ) }}
					{{ $fieldInline  = false }}
				{{ else }}
					{{ sendMessage nil "Accepted values for argument 4, <Value>, when using `fieldInline`: `\"true\"`, `\"false\"`." }}
                {{ end }}
			{{ end }}

			{{/* Update field to be changed in fields slice */}}
			{{ $fields.Set $fieldIndex (sdict "name" $fieldName "value" $fieldValue "inline" $fieldInline ) }}

		{{ else }}
			test3
		{{ end }}
	{{ else }}
		{{ sendMessage nil "Accepted values for argument 3, <Subcommand>: `author`, `thumbnail`, `title`, `description`, `fieldName`, `fieldValue`, `fieldInline`, `footer`." }}
	{{ end }}

	{{/* Build new embed */}}
	{{/* Define temp embed as old embed */}}
	{{ $embedUpdate := $embedOld }}

    {{/* Generate current timestamp for update */}}
    {{ $timestamp   := str (print currentTime.Year "-" (toInt currentTime.Month) "-" currentTime.Day "T" currentTime.Hour ":" currentTime.Minute ":" currentTime.Second "+00:00") }}

	{{ if (gt $numFields 0 ) }}
		{{ $embedUpdate = cembed
			"color" 000000
			"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
			"thumbnail" (sdict "url" $embedThumbnail )
			"title" ($embedTitle)
			"description" $embedDescription
			"fields" $fields  
            "footer" $embedFooter
            "timestamp" $timestamp
        }} 
	{{ else }}
		{{ $embedUpdate = cembed
			"color" 000000
			"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
			"thumbnail" (sdict "url" $embedThumbnail )
			"title" ($embedTitle)
			"description" $embedDescription
            "footer" $embedFooter
            "timestamp" $timestamp
		}} 
	{{ end }}

	{{/* Output old embed to log */}}
	{{ sendMessage 987256217179131954 $embedOld }}

	{{/* Update original embed in place */}}
	{{ editMessage $embedChannelID $embedMessageID $embedUpdate }}

    {{/* Output old embed to log */}}
	{{ sendMessage nil $embedUpdate }}

	{{ else  }}

		{{ sendMessage nil "Too many arguments, enclose <value> in double quotes if you are not already doing so and try again." }}

	{{ end }}
