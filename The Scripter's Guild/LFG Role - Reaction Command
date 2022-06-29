{{ $reactionMessageID := .ReactionMessage.ID }}
{{ $reactionaryID := (userArg .User.ID).ID }}

{{ print $reactionaryID }}

{{ $reactionaryNickName := (reReplace "#.*" (str (userArg .User.ID)) "" ) }}
{{ $emoji := .Reaction.Emoji.Name }}
{{ $lfgRoleMessage := 987224454931292221 }}

{{ $lfgList     := (dbGet 1 "lfgList").Value }}
{{ $lfgListLen  := len $lfgList }}

{{ print $lfgListLen }}


{{ if ( eq $reactionMessageID $lfgRoleMessage ) }}

{{ if ( eq $emoji "✅" ) }}
    {{ addRoleID 991421417738477668 }}

    {{ if ( eq ( index $lfgList 0 ) "empty" ) }}
        
        {{ dbSet 1 "lfgList" ( $lfgList.Set 0 ( print "<@" ( str $reactionaryID ) ">" ) )  }}

    {{ else }}     

        {{ dbSet 1 "lfgList" ( $lfgList.Append ( print "<@" ( str $reactionaryID ) ">" ) )  }}

        {{ sendDM ( print   "You are now set as LFG!\n"
                        "This will last for 30 minutes.\n"
                        "You can click the Emoji again to reset your timer.\n"
                        "The following mebers are also set to LFG:\n"
                        $lfgList ) }}

    {{ end }}

    {{ execCC 10 nil 0 "" }}
{{ else }}    
    {{ sendDM "You used the wrong Emoji." }}
{{ end }}

{{ else }} {{ end }}