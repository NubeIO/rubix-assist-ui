import Vue from "vue";


export default Vue.extend({
    name: "Home",
    data: () => ({
        aa: 2,
    }),
    methods: {
        getLocale2() {
            this.aa++;
        },
    },
});
