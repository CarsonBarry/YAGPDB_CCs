{{  if
( eq ( index .ReactionMessage.Reactions 0).Emoji.Name "🚫" )
}}
    {{ deleteMessage nil .ReactionMessage.ID }}

{{ else if
(eq (index .ReactionMessage.Reactions 0).Emoji.Name "👥")
}}
	{{ $embed := (index .Message.Embeds 0) }}

    {{- /* Define emtpy variables for elements that might not be there to copy */ -}}
    {{ $authorName				:= "" }}
    {{ $authorURL				:= "" }}
    {{ $authorIcon				:= "" }}
	{{- /* Unpack Embed into variables */ -}}
    {{ if $embed.Author }}
        {{ $authorName				:= str $embed.Author.Name }}
        {{ $authorURL				:= str $embed.Author.URL }}
        {{ $authorIcon				:= str $embed.Author.IconURL }}
    {{ else }} {{ end }}
	{{ $embedTitle 				:= str $embed.Title }}
	{{ $embedDescription		:= str $embed.Description }}
	{{ $embedThumbnail 			:= str $embed.Thumbnail }}
	{{ $embedFooter 			:= $embed.Footer }}
    {{ $embedFooterText         := "" }}

	{{- /* Initialize/define variables */ -}}
	{{ $numFields := len $embed.Fields }}
	{{ $fields := cslice "temp" }}

	{{- /* If: there is at least 1 Field */ -}}
	{{ if ( gt $numFields 0 ) }}

		{{- /* Store the value of that field as index 0 of the \$fields slice */ -}}
		{{ $fields = cslice (index $embed.Fields 0) }}

		{{- /* For each: Field after the first */ -}}
		{{ $i := 1 }} {{ while ( lt $i $numFields ) }}
			{{- /* Append next field to $fields slice */ -}}
			{{ $fields = $fields.Append (index $embed.Fields $i) }}
			{{- /* Iterate Loop */ -}}
			{{ $i = ( add $i 1 ) }}
		{{ end }}

	{{ else }}
	{{ end }}

	{{- /* Build new embed */ -}}
	{{- /* Define temp embed as old embed */ -}}
	{{ $embedCopy := $embed }}

    {{- /* Generate current timestamp for update */ -}}
    {{ $month := ( toInt currentTime.Month ) }}
    {{ $day := ( currentTime.Day ) }}
    {{ $hour := ( currentTime.Hour ) }}
    {{ $minute := ( currentTime.Minute ) }}
    {{ if (lt $month 10) }}
        {{ $month = ( print "0" ( str $month ) ) }}
    {{ else }} {{ end }}
    {{ if (lt $day 10) }}
        {{ $day = ( print "0" ( str $day ) ) }}
    {{ else }} {{ end }}
    {{ if (lt $hour 10) }}
        {{ $hour = ( print "0" ( str $hour ) ) }}
    {{ else }} {{ end }}
    {{ if (lt $minute 10) }}
        {{ $minute = ( print "0" ( str $minute ) ) }}
    {{ else }} {{ end }}
    {{ $timestamp   := str (print currentTime.Year "-" $month "-" $day "T" $hour ":" $minute ":00+00:00") }}

	{{ if (gt $numFields 0 ) }}
		{{ $embedCopy = cembed
			"color" 000000
			"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
			"thumbnail" (sdict "url" $embedThumbnail )
			"title" ($embedTitle)
			"description" $embedDescription
			"fields" $fields  
            "footer" $embedFooter
            "timestamp" ($timestamp)
        }} 
	{{ else }}
		{{ $embedCopy = cembed
			"color" 000000
			"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
			"thumbnail" (sdict "url" $embedThumbnail )
			"title" ($embedTitle)
			"description" $embedDescription
            "footer" $embedFooter
            "timestamp" ($timestamp)
		}} 
	{{ end }}

    {{- /* Send copy of embed to same channel*/ -}}
	{{ $messageID := sendMessageRetID nil $embedCopy }}

    {{- /* Take the channel ID from that message object and use the channel ID and message ID to generate a new footer for the new embed */ -}}
    {{ $channelID := ( getMessage nil $messageID ).ChannelID }}
    {{ $embedFooterText = ( print "chID: " (str $channelID) " msgID: " (str $messageID) ) }}
    {{ $embedFooter = (sdict "text" $embedFooterText) }}

    {{- /* Build the new embed */ -}}
    {{ if (gt $numFields 0 ) }}
		{{ $embedCopy = cembed
			"color" 000000
			"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
			"thumbnail" (sdict "url" $embedThumbnail )
			"title" ($embedTitle)
			"description" $embedDescription
			"fields" $fields  
            "footer" $embedFooter
            "timestamp" ($timestamp)
        }} 
	{{ else }}
		{{ $embedCopy = cembed
			"color" 000000
			"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
			"thumbnail" (sdict "url" $embedThumbnail )
			"title" ($embedTitle)
			"description" $embedDescription
            "footer" $embedFooter
            "timestamp" ($timestamp)
		}} 
	{{ end }}

    {{- /* Update the embed in place with the new footer */ -}}
    {{ editMessage $channelID $messageID $embedCopy }}

    {{ deleteMessage nil .ReactionMessage.ID }}

{{ else }}
  {{ sendMessage nil "Dupe/Delete ran, even though you used a different Emoji. Nothing happened." }}
{{ end }}

