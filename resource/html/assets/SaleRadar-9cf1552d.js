import{du as r}from"./useECharts-819d58cc.js";import{N as s}from"./antd-12f11a56.js";import{d as n,f as i,w as d,a8 as l,_ as m,a9 as f,aa as p,a0 as u,ag as c}from"./vue-5c68ae35.js";import{_ as h}from"./index.js";const g=n({components:{Card:s},props:{loading:Boolean,width:{type:String,default:"100%"},height:{type:String,default:"400px"}},setup(e){const a=i(null),{setOptions:t}=r(a);return d(()=>e.loading,()=>{e.loading||t({legend:{bottom:0,data:["Visits","Sales"]},tooltip:{},radar:{radius:"60%",splitNumber:8,indicator:[{name:"2017"},{name:"2017"},{name:"2018"},{name:"2019"},{name:"2020"},{name:"2021"}]},series:[{type:"radar",symbolSize:0,areaStyle:{shadowBlur:0,shadowColor:"rgba(0,0,0,.2)",shadowOffsetX:0,shadowOffsetY:10,opacity:1},data:[{value:[90,50,86,40,50,20],name:"Visits",itemStyle:{color:"#9f8ed7"}},{value:[70,75,70,76,20,85],name:"Sales",itemStyle:{color:"#1edec5"}}]}]})},{immediate:!0}),{chartRef:a}}});function w(e,a,t,y,S,_){const o=l("Card");return m(),f(o,{title:"销售统计",loading:e.loading},{default:p(()=>[u("div",{ref:"chartRef",style:c({width:e.width,height:e.height})},null,4)]),_:1},8,["loading"])}const $=h(g,[["render",w]]);export{$ as default};
