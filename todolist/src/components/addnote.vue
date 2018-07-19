<template>
    <div class="logtext1">
        <div style="line-heigh:82px;font-size:28px;">NEW NOTE</div>
        <br />
        <div class="logtext1d">
            <span>event name</span>
            <input class="input1" type="text" v-model="evname"/>
        </div>
        <div class="logtext2d">
            <span>start time</span>
            <input class="input2" type="date" value="" v-model="statime"/>
        </div>
        <br />
        <button class="btncr" v-on:click="cre_note">create</button>
        <br />
        <br />
        <br />
    </div>
</template>

<script>
export default {
   data(){
       return{
            evname:'',
            statime:'',
            noteinfo:
                {
                    label:'',
                    time:''
                }
            
       }
   },
   methods:{
       cre_note:function(){
           this.noteinfo.label=this.evname
           this.noteinfo.time=this.statime
           this.$emit('getNoteInfo', this.noteinfo)
           this.evname=''
           this.statime=''
       }
   }
    
}
</script>

<style>
    .logtext1{
        max-width:500px;
        margin-top:10%;
        margin-left:33%;
        background-color:#6b6260;
        position:relative;
        z-index:10;
        border-radius: 0.5em;
        color:lemonchiffon;
        transition:ease 1s;
    }
    .logtext1d{
        font-size:18px;
    }
    .logtext2d{
        font-size:18px;
        margin-top:15px;
        margin-left:-24%;
    }
    .input1{
        font-size:17px;
        width:300px;
        height:23px;
    }
    .input2{
        height:23px;
        font-size:17px;
    }
    .btncr{
        width:90px;
        height:30px;
        font-size:20px;
        color:lavender;
        background-color: #6b6260;
        border:0.5px solid;
    }
    
</style>