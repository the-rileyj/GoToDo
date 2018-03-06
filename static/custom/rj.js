var v = new Vue({
    el: '#app',
    data: function() {
        return {
            date: "",
            eventList: []
        }
    },
    methods: {
        selectedDate: function(date) {
            
        }
    },
    mounted: function() {
        // v = this
        // axios.get(window.location.href + "api/creds/nasa")
        // .then((res) => {
        //     v.key = res.data.key
        // })
        // .catch((err) => {
        //     //Though not nessesary with how the webserver is structured, this is hypothetically where error handling would go
        // });
    }
});