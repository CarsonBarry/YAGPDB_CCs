it ran
{{/* Pull project counts from database */}}
{{ $infiniteProjectCount 	:= toInt (dbGet 1 "infiniteProjectCount").Value }}
{{ $legacyHaloProjectCount 	:= toInt (dbGet 1 "legacyHaloProjectCount").Value }}
{{ $unrealProjectCount 		:= toInt (dbGet 1 "unrealProjectCount").Value }}
{{ $unityProjectCount 		:= toInt (dbGet 1 "unityProjectCount").Value }}
{{ $coreProjectCount 		:= toInt (dbGet 1 "coreProjectCount").Value }}
{{ $fortniteProjectCount 	:= toInt (dbGet 1 "fortniteProjectCount").Value }}

{{/* Define static values */}}
{{ $authorText := "TSG Admin Team" }}
{{ $authorURL := "https://discord.com/channels/220766496635224065/746081416940617878" }}
{{ $authorIcon := "https://i.imgur.com/505zjJl.png" }}
{{ $thumbnail := "https://i.imgur.com/505zjJl.png" }}

{{/* Get variable targeting for reactionMessage */}}
{{ $reactionChannelID := .ReactionMessage.ChannelID }}
{{ $reactionMessageID := .ReactionMessage.ID }}

{{ if and
	(eq (index .ReactionMessage.Reactions 0).Emoji.Name "v")
	(eq (index .ReactionMessage.Reactions 0).Count 2)
}}
	{{ $embedInput := (index .Message.Embeds 0) }}
	
	{{ $quickSwap := sendMessageRetID nil $embedInput }}
	
	{{ $embedInput := (index (getMessage nil $quickSwap).Embeds 0) }}

	{{ $editorfield := (index $embedInput.Fields 0).Value }}
	{{ $title := (index $embedInput.Fields 1).Value }}
	{{ $description := (index $embedInput.Fields 2).Value }}

	{{/* Process User ID */}}
	{{ $embedFooter 			:= $embedInput.Footer }}
    {{ $embedFooterText         := $embedFooter.Text }}
	{{ $projectOriginator		:= ( reReplace "User ID: " $embedFooterText "" ) }}
		
	{{ $messageTarget := "nil" }}
	{{ $projectCount := "0" }}
	{{ $titlePrefixCode := "ZZ" }}

	{{ if (reFind `:dynoSuccess: Legacy Halo` $editorfield) }}
		{{ $messageTarget = "952262627084607508" }}
		{{ $projectCount = $legacyHaloProjectCount }}
		{{ $titlePrefixCode = "LH" }}
	{{ else if (reFind `:dynoSuccess: Halo` $editorfield) }}
		{{ $messageTarget = "952261623928742028" }}
		{{ $projectCount = $infiniteProjectCount }}
		{{ $titlePrefixCode = "HI" }}
	{{ else if (reFind `:dynoSuccess: Unreal` $editorfield) }}
		{{ $messageTarget = "952261686121869382" }}
		{{ $projectCount = $unrealProjectCount }}
		{{ $titlePrefixCode = "UE" }}
	{{ else if (reFind `:dynoSuccess: Unity` $editorfield) }}
		{{ $messageTarget = "952261778287517766" }}
		{{ $projectCount = $unityProjectCount }}
		{{ $titlePrefixCode = "UN" }}
	{{ else if (reFind `:dynoSuccess: Core` $editorfield) }}
		{{ $messageTarget = "952261728983482368" }}
		{{ $projectCount = $coreProjectCount }}
		{{ $titlePrefixCode = "CO" }}
	{{ else if (reFind `:dynoSuccess: Fortnite` $editorfield) }}
		{{ $messageTarget = "952261829411889242" }}
		{{ $projectCount = $fortniteProjectCount }}
		{{ $titlePrefixCode = "FN" }}
	{{ else if (reFind `:dynoSuccess: Other` $editorfield) }}
		{{ $messageTarget = "982661501766754355" }}
		{{ $projectCount = "0" }}
		{{ $titlePrefixCode = "MP" }}
	{{ end }}
	
	{{ if (lt $projectCount 10 ) }}
	
		{{ $projectCount = print "0" $projectCount }}
	{{ else }}  {{ end }} 
	{{ $title = print $titlePrefixCode "-" $projectCount }}
	
	{{ $embedOutput := cembed
		"author" ( sdict "name" $authorText "url" $authorURL "icon_url" $authorIcon ) 
		"title" ( $title ) 
		"description" ( $description ) 
		"color" 000000
		"thumbnail" ( sdict "url" $thumbnail ) 
		"fields" ( cslice
			( sdict "name" "Links" "value" "-" "inline" false ) 
			( sdict "name" "Additional Info" "value" "-" "inline" false ) 
			( sdict "name" "Version" "value" "0.0.1" "inline" true ) 
			( sdict "name" "Project Leads" "value" "-" "inline" true ) 
			( sdict "name" "Contributors" "value" "-" "inline" true ) 						
		) 
		"footer" ( sdict "text" $projectOriginator ) 
	}} 
	
	{{/* Output embed to log */}}
	{{ sendMessage 987256217179131954 $embedOutput }}
	

	{{/* Send embed to Project Channel */}}
	{{ sendMessage $messageTarget $embedOutput }}

	{{/* Cleanup */}}
	{{ deleteMessage nil $quickSwap }}
	{{ deleteAllMessageReactions $reactionChannelID	$reactionMessageID }}
	{{ addReactions "✅" }}

{{ else if and
(or
	(eq (index .ReactionMessage.Reactions 0).Emoji.Name "✅")
	(eq (index .ReactionMessage.Reactions 0).Emoji.Name "❎")
)
(eq (index .ReactionMessage.Reactions 0).Count 2) 
}}
 
{{ deleteAllMessageReactions $reactionChannelID	$reactionMessageID }}
{{ addReactions "👍" "👎" }}

{{ else if and
(eq (index .ReactionMessage.Reactions 1).Emoji.Name "👎")
(eq (index .ReactionMessage.Reactions 1).Count 2)
}}
	{{ deleteAllMessageReactions $reactionChannelID	$reactionMessageID }}
	{{ addReactions "❎" }}

{{else}} You did it wrong. Ask Cap for help. {{end}}
