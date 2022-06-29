{{/* Run this command to set up your LFG List Database Entry */}}

{{ dbSet 1 "lfgList" ( cslice "empty" ) }}

{{ $lfgList     := (dbGet 1 "lfgList").Value }}
{{ $lfgListLen  := len $lfgList }}

{{ print $lfgList }}