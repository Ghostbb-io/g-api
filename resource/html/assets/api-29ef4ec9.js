import{d as e}from"./index.js";const i=t=>e.get({url:"/v1/api/page",params:t}),s=t=>e.post({url:"/v1/api",params:t}),r=(t,p)=>e.put({url:`/v1/api/${t}`,params:p}),o=t=>e.delete({url:`/v1/api/${t}`}),d=()=>e.get({url:"/v1/api/tree"});export{s as a,d as b,o as d,r as e,i as g};