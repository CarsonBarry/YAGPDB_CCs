{{ $reactionChannelID := .ReactionMessage.ChannelID }}
{{ $reactionMessageID := .ReactionMessage.ID }}
{{ $reactionMessageA uthor := str .ReactionMessage.Author.ID }}
{{ $reactionary := (reReplace "#.*" (str (userArg .User.ID)) "" ) }}
{{ $reactionChannelName := (getThread .ReactionMessage.ChannelID).Name }}
{{ $emoji := .Reaction.Emoji.Name }}

{{if ( eq $reactionMessageAuthor "878399831238909952" ) }}

    {{ if (reFind $reactionary $reactionChannelName) }}
        {{ if (eq $emoji "✅") }}
            {{ deleteMessage $reactionChannelID $reactionMessageID }}
        {{ else }}    
            {{ sendDM "You used the wrong Emoji." }}
        {{ end }}
    {{ else }}   
        {{ if (eq $emoji "✅") }}
            {{ sendDM "You do not have permissions to delete this message." }}
        {{ else }} 
            {{ sendDM "You used the wrong Emoji *and* you do not have permissions to delete this message." }}
        {{ end }}
    {{ end }}
{{ else }}

{{ end }}