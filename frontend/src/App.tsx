import { Route, Routes } from "react-router-dom";

import IndexPage from "@/pages/index";
import OmciAnalyzerPage from "@/pages/omcianalyzer/index";
import OmciAnalyzerPageOnusPage from "@/pages/omcianalyzer/onus";
import OmciAnalyzerDiagramPage from "@/pages/omcianalyzer/diagram";
import OmciAnalyzerPageOmciPage from "@/pages/omcianalyzer/omci";
import LibraryPage from "@/pages/library/index";
import ConfigAnalyzerPage from "@/pages/configanalyzer/index";
import SequenceTracerPage from "@/pages/sequencetracer/index";
import DownloadsPage from "@/pages/downloads/index";
import AboutPage from "@/pages/about";
import CollectorPage from "@/pages/collector/index";
import CollectorLogsPage from "./pages/collector/logs";

function App() {
  return (
    <Routes>
      <Route element={<IndexPage />} path="/index" />
      <Route element={<OmciAnalyzerPage />} path="/" />
      <Route element={<OmciAnalyzerPage />} path="/omcianalyzer" />
      <Route element={<OmciAnalyzerPageOnusPage />} path="/omcianalyzer/onus" />
      <Route element={<OmciAnalyzerDiagramPage />} path="/omcianalyzer/diagram" />
      <Route element={<OmciAnalyzerPageOmciPage />} path="/omcianalyzer/omci" />
      <Route element={<LibraryPage />} path="/library" />
      <Route element={<CollectorPage />} path="/collector" />
      <Route element={<CollectorLogsPage />} path="/collector/logs" />
      <Route element={<ConfigAnalyzerPage />} path="/configanalyzer" />
      <Route element={<SequenceTracerPage />} path="/sequencetracer" />
      <Route element={<DownloadsPage />} path="/downloads" />
      <Route element={<AboutPage />} path="/about" />
    </Routes>
  );
}

export default App;
