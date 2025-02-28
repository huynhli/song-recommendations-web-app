import { useEffect, useState } from 'react'
import React from 'react';
import { BrowserRouter, Routes, Route, Outlet } from 'react-router-dom';
import './App.css'
import Header from './Header'
import Footer from './Footer'
import GeneratorPage from './GeneratorPage';
import DocsPage from './DocsPage';

export default function App() {
  // defining default layout
  const Layout = () => {
    return (
      <div>
        <Header/>
        <Outlet/>
        <Footer/>
      </div>
    )
  }

  return (
    <div>
      <Routes>
        <Route path="/" element={<Layout />}>
            <Route path="/" element={<GeneratorPage />} />
            <Route path="/docs" element={<DocsPage />} />
        </Route>
      </Routes>
    </div>
  );
}
