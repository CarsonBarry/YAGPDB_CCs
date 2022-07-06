# YAGPDB_CCs
YAGPDB CC's I've written to handle various tasks for the servers I admin.

`Servers`

**The Scripter's Guild** - 
*UGC Collab and Learning. Halo Forge/Modding Focused, but also branching into Unreal, Core, Unity, FNC, and related tools.*

**ForgeHub** - 
*Halo's OG UGC Forum and content directory.*

**The Foundry** - 
*A group for established Halo content creators.*

**Biggy Boi's Cave** - 
*The personal server of streamer Sir Biggy Mac.*

`Commands`

**LFG Role**

Basic
*Creates an embed and places the <inserthere> emoji as a reaction on the embed's message. Users who react to that message with the same reaction will have your LFG role assigned and a 30 minute delay will be set to remove it. Users will receive a direct message letting them know they are set to LFG. Users can only run this command if they are not currently set to LFG.*

Advanced
*Same as Basic, but adds more emoji reactions/functions and also stores the current LFG roster in fields in the embed for different topics. Roster updates every 5 minutes. LFG role status is controlled via information changes in embed rather than delayed removal right after adding the role; this allows users to 're-up' and reset their timer before it runs out as well as remove themselves from the LFG roster early.*

**Embed Manipulation Commands**
react: Quick Dupe/Delete
-embed 

**Project Channel System** - *Req: Dyno, Bot (Forms)*
  `-projectCount`
  `react: Form Approval`
  `-project`

**Thread Share System** - *Req: Needle, Bot (Autothreading)*
  `-shareSetup`
  `-share`

**Configurable Search Link Generator**
  `-search`
