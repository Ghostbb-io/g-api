import{f as r,_ as o}from"./index.js";import{b as p}from"./index-086c3175.js";import{d as m,a8 as i,_ as u,$ as l,a0 as c,a1 as _,l as d,ad as f,a2 as b}from"./vue-5c68ae35.js";import{ai as g}from"./antd-12f11a56.js";import"./index-ae722a40.js";import"./index-2bcf2ae6.js";import"./index-3f5dd244.js";import"./useContentViewHeight-c2d18fda.js";import"./useWindowSizeFn-997fa1d0.js";import"./lock-56c7862a.js";const C=m({name:"InputNumberItem",components:{InputNumber:g},props:{event:{type:Number},title:{type:String}},setup(e){const{prefixCls:t}=r("setting-input-number-item");function n(s){e.event&&p(e.event,s)}return{prefixCls:t,handleChange:n}}});function I(e,t,n,s,N,v){const a=i("InputNumber");return u(),l("div",{class:b(e.prefixCls)},[c("span",null,_(e.title),1),d(a,f(e.$attrs,{size:"small",class:`${e.prefixCls}-input-number`,onChange:e.handleChange}),null,16,["class","onChange"])],2)}const E=o(C,[["render",I],["__scopeId","data-v-7ccf252c"]]);export{E as default};