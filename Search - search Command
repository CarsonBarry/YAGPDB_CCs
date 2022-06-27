{{/* Define Locations */}}
{{ $searchLocation	:= cslice	"Proper Name" "locationArgument" "Root URL" "Icon URL"
								"Search URL Sabs Terms" }}
{{ $forgehubForum	:= cslice	"Forgehub Forum" "fhforum" "https://www.forgehub.com/forum" ""
								"https://www.google.com/search?q=site:https://www.forgehub.com/forums+" }}
{{ $forgehubWiki	:= cslice	"Forgehub Wiki" "fhwiki" "https://www.forgehub.com/wiki" ""
								"https://www.google.com/search?q=site:https://www.forgehub.com/wiki+" }}

{{ if (eq (len .CmdArgs) 1) }}
	
	{{ $args := parseArgs 1 "Syntax is -embedUpdate <help>"
		(carg "string" "Accepts 'help' if only argument given.")
	}}

		{{ if (eq ($args.Get 0) "help") }}

		{{ $helpMessage := cembed
			"color" 000000
			"author" (sdict	"name" ("TSG Admin Team") "url" ("https://discord.gg/ZpE3tcpZy9") "icon_url" ("https://i.imgur.com/505zjJl.png"))
			"thumbnail" (sdict "url" ("https://i.imgur.com/a7NAJZU.png") )
			"title" ("Search Command Help")
			"description" (print
				"*Use this command to generate a link to a search of the desired location.*\n\n"
				"**Note:** Once you are there, please use the search at that location instead of continuously generating more searches with this command.\n\n"
				"Syntax:\n\n"
				"**1 Argument Given**\n"
				"`-search <help>`\n*displays this message*\n\n"
				"**2 Arguments Given**\n"
				"`-search <location> <searchTerms>`\n*Generates search url using given terms and location.*\n\n"
				"**location** - *The site you are trying to generate a search for.*\n\n"
				"**searchTerms** - *The terms you are trying to search. Enclose this argument in double quotes for a more consistent experience.*\n\n"
			)
		}}

		{{ sendMessage nil $helpMessage }}
	
		{{ else }}

			{{ sendMessage nil "Use `-search help` to see the full syntax and an explanation of this command."}}

		{{ end }}
	
{{ else if (eq (len .CmdArgs) 2) }}

	{{ $args := parseArgs 2 "Syntax is -search <location> <searchTerms>"
		(carg "string" "Location to search. Accepted values: ")
		(carg "string" "The terms you would like to search. Enclose this argument in double quotes for a more consistent experience.")
	}}

	{{/* Package args into variables */}}
	{{ $location	 	:= ( $args.Get 0 ) }}
	{{ $searchTerms		:= ( $args.Get 1 ) }}

	{{ if ( eq $location "fhforum" ) }}

		{{ $searchLocation = $forgehubForum }}

	{{ else if ( eq $location "fhwiki" ) }}

		{{ $searchLocation = $forgehubWiki }}

	{{ else }}

		{{ sendMessage nil "Accepted values for the `location` argument: `fhforum`, `fhwiki`." }}

	{{ end }}

	{{ $location		:= index $searchLocation 0 }}
	{{ $locationRootURL := index $searchLocation 2 }}
	{{ $locationIcon	:= index $searchLocation 3 }}
	{{ $searchURLBase	:= index $searchLocation 4 }}
	{{ $search			:= (print "Search \"" $searchTerms "\" on the " $location ) }}
	{{ $searchTerms		 = (reReplace " " $searchTerms "+") }}
	{{ $results		    := (print $searchURLBase $searchTerms ) }}

	{{/* Generate timestamp */}}
	{{ $timestamp	:= str (print currentTime.Year "-" (toInt currentTime.Month) "-" currentTime.Day "T" currentTime.Hour ":" currentTime.Minute ":" currentTime.Second "+00:00") }}

	{{ $embed := cembed
			"color" 000000
			"author" ( sdict "name" ($location) "url" ($locationRootURL) "icon_url" ("") )
			"title" $search
			"description" $results
			"footer" ( sdict "text" "Search Results Link ^^" )
	}} 
	
	{{ sendMessage nil $embed }}

{{ else }}

	{{ sendMessage nil "Syntax is `-search <location> <searchTerms>` or `-search <help>`" }}

{{ end }}
