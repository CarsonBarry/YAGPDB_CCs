{{ $lfgList     := (dbGet 1 "lfgList").Value }}
{{ $lfgListLen  := len $lfgList }}
{{ $name 		:= "temp" }}
{{ $user 		:= "temp" }}

{{ $lfgListFormatted := "" }}

{{ range $key, $value := $lfgList }}
	{{- $name = $value -}}
	{{- $lfgListFormatted = ( print $lfgListFormatted "\n" $value ) -}}
{{ end }}

{{/* Define Embed Content */}}
{{ $authorName			:= "TSG Admin Team" }}
{{ $authorURL			:= "" }}
{{ $authorIcon			:= "" }}
{{ $embedTitle 			:= "LFG" }}
{{ $embedDescription	:= ( print	"**Click the reaction to add yourself to the list!**\n"
									"*This will last for 30 minutes at a time, you can always reapply the role.*\n\n"
									"`Members that are currently LFG (Ping them to play!)`"
									$lfgListFormatted ) }}
{{ $embedThumbnail 		:= "" }}
{{ $embedFooterText     := "" }}
{{ $embedFooter 		:= (sdict "text" $embedFooterText) }}

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


{{ $embed := cembed
	"color" 000000
	"author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
	"thumbnail" (sdict "url" $embedThumbnail )
	"title" ($embedTitle)
	"description" $embedDescription
	"footer" $embedFooter
	"timestamp" $timestamp
}}

{{ editMessage 929992569981640744 987224454931292221 $embed }}
