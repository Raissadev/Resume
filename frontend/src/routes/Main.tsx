import React from "react";
import { Routes, Route } from "react-router-dom";

import Head from "../components/Header";
import Foot from "../components/Footer";
import "../styles/App.less"

import Home from "../views/Home"

export default function Main(): any
{
    return(
        <>
            <Head />
                <Routes>
                    <Route path="/" element={ <Home /> } />
                </Routes>
            <Foot />
        </>
    )
}