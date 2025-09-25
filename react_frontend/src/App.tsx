import { useEffect, useState } from 'react'
import React from 'react';
import { BrowserRouter, Routes, Route, Outlet } from 'react-router-dom';
import './App.css'
import Header from './pages_or_components/Header'
import Footer from './pages_or_components/Footer'
import GeneratorPage from './pages_or_components/GeneratorPage';
import DocsPage from './pages_or_components/DocsPage';
import HomePage from './pages_or_components/HomePage';

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
            <Route path="/" element={<HomePage />} />
            <Route path="/recsGenerator" element={<GeneratorPage />} />
            <Route path="/docs" element={<DocsPage />} />
        </Route>
      </Routes>
    </div>
  );
}
