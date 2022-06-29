

{{ if not   (or 
                (eq (len .CmdArgs) 1) 
                (eq (len .CmdArgs) 3) 
            ) }}

    {{ sendMessage nil "Syntax is `-project <embedMessageID> <subCommand> <value>` or `-project help`"}}

{{ else if (eq (len .CmdArgs) 1) }}

        {{ $args := parseArgs 1 "Syntax is -project <help>"
            (carg "string" "Accepts 'help' if only argument given.")
        }}

        {{ if (eq ($args.Get 0) "help") }}

            {{ $helpMessage := cembed
                "color" 000000
                "author" (sdict	"name" ("TSG Admin Team") "url" ("https://discord.gg/ZpE3tcpZy9") "icon_url" ("https://i.imgur.com/505zjJl.png"))
                "thumbnail" (sdict "url" ("https://i.imgur.com/a7NAJZU.png") )
                "title" ("Project Command Help")
                "description" (print
                    "*Use this command to update/edit your Project Embed. You have to properly target the embed and it has to be done from the relevant thread for the commands to work. You must be a lead on the relvant project or a Guild Officer to use this command.*\n\n"
                    "Syntax:\n\n"
                    "`-project <help>`\n*displays this message*\n\n"
                    "`-project <embedMessageID> <subCommand> <value>`\nPerforms given subcommand on targeted embed with given value.\n\n"
                    "**embedMessageID** - *Embed Message ID. Copy this from the footer of the Project Embed. Also accepts 'help' if only argument given.*\n\n"
                    "**subCommand** - *Subcommand. Accepted values: version, name, description, links, addLead, removeLead*\n\n"
                    "**value** - *Relevant input per Subcommand. 'description' and 'links' subcommands accept normal message formatting and \\n as a line break, but you can just add line breaks in your command message.*"
                )
            }}

            {{ sendMessage nil $helpMessage}}

        {{ else }}

            {{ sendMessage nil "Use `-project help` to see this command's full syntax and an explanation."}}

        {{ end }}

{{ else }}
    {{/* Can only get here if there are 3 args */}}

    {{ $args := parseArgs 3 "Syntax is -project <embedMessageID> <subCommand> <value>"
        (carg "string" "Embed Message ID. Copy this from the footer of the Project Embed. Also accepts 'help' if only argument given.")
        (carg "string" "Subcommand. Accepted values: version, name, description, links, addLead, removeLead")
        (carg "string" "Relevant input per Subcommand. 'description' and 'links' subcommands accept normal message formatting and \\n as a line break, but you can just add line breaks in your command message.")
    }}

    {{/* Package Args into Variables */}}
    {{ $embedMessageID	    		:= ($args.Get 0) }}
    {{ $subCommand 	        		:= ($args.Get 1) }}
    {{ $value           	    	:= ($args.Get 2) }}

    {{/* Is Member a Mod or a Lead on this Project? */}} 
    {{ $userID 	                    := str .User.ID }}
    {{ $embedChannelID 	        	:= .Channel.ParentID }}     
    {{ $embedOld                	:= (index (getMessage $embedChannelID $embedMessageID).Embeds 0) }}
    {{ $projectLeads            	:= str (index $embedOld.Fields 1).Value  }} 

    {{ if or (reFind $userID $projectLeads) (hasRoleID 735183135729385542) }}

        {{/* Unpack Old Embed - (minus possible changing fields) */}}

        {{ $authorName	    		:= str $embedOld.Author.Name }}
        {{ $authorURL		    	:= str $embedOld.Author.URL }}
        {{ $authorIcon	    		:= str $embedOld.Author.IconURL }}
        {{ $embedThumbnail 	    	:= str $embedOld.Thumbnail }}
        {{ $messageID		    	:= str (index $embedOld.Fields 2).Value }}

        {{/* Define Dynamic Field Variables */}}
        {{ $embedTitle 	    		:= str $embedOld.Title }}
        {{ $embedDescription    	:= str $embedOld.Description }}
        {{ $projectVersion 	    	:= str (index $embedOld.Fields 0).Value }}
        {{ $projectLeads        	:= str (index $embedOld.Fields 1).Value  }} 

        {{ if (eq $subCommand "name")}}
            {{ $embedTitle 	    	= $value }}
        {{ else }}       
        {{ end }}

        {{ if (eq $subCommand "description")}}
            {{ $embedDescription	= $value }}
        {{ else }}    
        {{ end }}

        {{ if (eq $subCommand "version")}}
            {{ $projectVersion 		= $value }}
        {{ else }}          
        {{ end }}

        {{ if (eq $subCommand "addLead")}}
            {{ $projectLeadsUpdate 	:= (print $projectLeads "\n<@" $value ">") }}
            {{ $projectLeads       	= $projectLeadsUpdate }}
            {{ sendMessage nil (print "<@" $value "> is now a project lead for this project!") }}
        {{ else if (eq $subCommand "removeLead")}}
            {{ $projectLeadsUpdate 	= (reReplace "\n\n" (reReplace (print "<@" $value ">") $projectLeads "") "") }}
            {{ $projectLeads    	= $projectLeadsUpdate }}
        {{ else }}
        {{ end }}

        {{/* Build new embed */}}
        {{ $embedUpdate := cembed
            "color" 000000
            "author" (sdict	"name" ($authorName) "url" $authorURL "icon_url" $authorIcon)
            "thumbnail" (sdict "url" $embedThumbnail )
            "title" ($embedTitle)
            "description" $embedDescription
            "fields" (cslice 
                (sdict "name" "Version" "value" $projectVersion "inline" true)
                (sdict "name" "Project Lead(s)" "value" $projectLeads "inline" true)
                (sdict "name" "Message ID" "value" (str $messageID) "inline" false)
        )
        }} 

        {{/* Output old embed to log */}}
        {{ sendMessage 987256217179131954 $embedOld }}

        {{/* Output new embed to thread */}}
        {{ sendMessage nil $embedUpdate }}

        {{/* Update original embed in place */}}
        {{ editMessage $embedChannelID $embedMessageID $embedUpdate }}

    {{else}}  

        {{ sendMessage nil "You are not a Lead on this project or a Guild Officer. Please seek assistance with the change you are trying to make." }}
        {{ sendMessage 987256217179131954 (print "<@" $userID "> tried using `-project " $subCommand "` on " $embedMessageID " with the value\n\n" $value )}}
   
    {{end}}

{{ end }}
