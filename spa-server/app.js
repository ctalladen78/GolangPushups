
// https://vuejs-tips.github.io/cheatsheet/
(function(Vue){
    let vm = new Vue({
        el: 'app',
        // mounted: function(){}
        created: function(){
            console.log("hello VUE 2")
        },
        data: {
            msg: "initial message"
        }

    })
    console.log("vm.$isServer ?", vm.$isServer)
})(Vue)