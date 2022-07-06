{{/*  */}}
{{ $update			:= true }}
{{ $new 			:= false }}
{{ $logChannel		:= "nil" }}
{{ $backupChannelID	:= "nil" }}
{{ $backupMessageID	:= "" }}
{{ $targetChannel 	:= "nil" }}
{{ $index 			:= 0 }}
{{ $value			:= "" }}
{{ $logMessage      := "Placeholder" }}

{{/*  */}}
{{ if not (eq (len .CmdArgs) 5) }}

	{{/* send help message */}} 
	{{ sendMessage nil "help message" }}

{{ else }}
	
	{{/* parse args */}}
	{{ $args := parseArgs 5 "Syntax is -embedUpdate <embedChannelID> <embedMessageID> <subCommand> <fieldIndex> <value>"
		(carg "string" "Channel ID of the channel where the message is. Copy this from the footer of the Project Embed.")	
		(carg "string" "Message ID of the message to which the embed is attached. Copy this from the footer of the Project Embed.")
		(carg "string" "Subcommand. Accepted values: `author`, `thumbnail`, `title`, `description`, `field`, `footer`.\nUsing `field` induces the requirement to provide a `fieldIndex` prior to the `value`, while not using `field` requires `nil` to be used as the `fieldIndex`.")
		(carg "string" "The index of the field that you are trying to update. Use `nil` if you are not editing fields.  Starts at 0 and ascends top to bottom (left to right when inline fields are involved).")
		(carg "string" "Relevant input per Subcommand. 'description' and 'links' subcommands accept normal message formatting and \\n as a line break, but you can just add line breaks in your command message. Enclose this argument in double quotes for a more consistent experience.")
	}}  

	{{/* fill variables */}}	
	{{ $embedChannelID		:= ($args.Get 0) }}
	{{ $embedMessageID		:= ($args.Get 1) }}
	{{ $subCommand			:= ($args.Get 2) }}
	{{ $arg4	 			:= ($args.Get 3) }}
	{{ $arg5				:= ($args.Get 4) }}
	{{ $embedOld			:= (index (getMessage $embedChannelID $embedMessageID).Embeds 0) }}
	{{ $authorName			:= str $embedOld.Author.Name }}
	{{ $authorURL			:= str $embedOld.Author.URL }}
	{{ $authorIcon			:= str $embedOld.Author.IconURL }}
	{{ $embedTitle			:= str $embedOld.Title }}
	{{ $embedDescription	:= str $embedOld.Description }}
	{{ $embedThumbnail		:= str $embedOld.Thumbnail }}
	{{ $embedFooter			:= $embedOld.Footer }}
	{{ $embedFooterText		:= $embedFooter.Text }}
	{{ $numFields			:= len $embedOld.Fields }}
	{{ $fields				:= cslice (sdict "name" "" "value" "" "inline" false ) }}
	
	{{/* process embed fields */}}
	{{ if ( gt $numFields 0 ) }}
		{{ $fields = cslice (index $embedOld.Fields 0) }}
		{{ $i := 1 }} {{ while ( lt $i $numFields ) }}
			{{- $fields = $fields.Append (index $embedOld.Fields $i) -}}
			{{- $i = ( add $i 1 ) -}}
		{{ end }}
	{{ end }}
	
	{{/* generate current timestamp in proper formatting */}}
	{{ $month := ( toInt currentTime.Month ) }} {{ if (lt $month 10 ) }} {{ $month = ( print "0" ( str $month ) ) }} {{ end }}
	{{ $day := ( currentTime.Day ) }} 			{{ if (lt $day 10 ) }} {{ $day = ( print "0" ( str $day ) ) }} {{ end }}
	{{ $hour := ( currentTime.Hour ) }} 		{{ if (lt $hour 10 ) }} {{ $hour = ( print "0" ( str $hour ) ) }} {{ end }}
	{{ $minute := ( currentTime.Minute ) }} 	{{ if (lt $minute 10 ) }} {{ $minute = ( print "0" ( str $minute ) ) }} {{ end }}
	{{ $timestamp := str ( print currentTime.Year "-" $month "-" $day "T" $hour ":" $minute ":00+00:00" ) }}
	
	{{/* define template to build embed and build embed inside embedUpdate variable as copy of old embed with new timestamp */}}
	{{ block "constructEmbed" }}
		{{- $embedUpdate = cembed
			"color" $color
			"author" (sdict	"name" $authorName
							"url" $authorURL
							"icon_url" $authorIcon )
			"thumbnail" ( sdict "url" $embedThumbnail )
			"title" $embedTitle
			"description" $embedDescription
			"fields" $fields
			"footer" $embedFooter
			"timestamp" $timestamp
		-}}
	{{ end }}

	{{/* TODO: add the remainder of subcommands to the if-else below and update syntax references above and logMessage constructor below */}}

	{{/* process subCommand */}}
	{{ if ( eq $subCommand "") }}
		{{ $ = $value }}
	{{ else if ( eq $subCommand "") }}
		{{ $ = $value }}
	{{ else }}
		{{/* send help message */}} 
		{{ sendMessage nil "help message" }}
	{{ end }}

	{{/* Output to log */}}
	{{ $logMessage := ( print .User.ID " ran `-project` with args:\n\n```" ($args.Get 0) "```\n\n```" ($args.Get 1) "```\n\n```" ($args.Get 2) "```\n\n```" ($args.Get 3) "```\n\n```" ($args.Get 4) "```" ) ) }}
	{{ sendMessage $logChannel $logMessage }}
	{{ sendMessage $logChannel $embedOld }}
	{{ sendMessage $logChannel $embedUpdate }}

	{{/* Update original embed in place */}}
	{{ if $update }} 
		{{ template "constructEmbed"
			$color $authorName $authorURL
			$authorIcon $embedThumbnail $embedTitle
			$embedDescription $fields $embedFooter
			$timestamp }}
		{{ editMessage $embedChannelID $embedMessageID $embedUpdate }}
	{{ end }}

	{{/* Send new embed to target location */}}
	{{ if $new }} {{ sendMessage $targetChannel $embedUpdate }} {{ end }}
{{ end }}