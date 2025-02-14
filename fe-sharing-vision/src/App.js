import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import CreatePost from './component/page/CreatePost';
import EditPost from './component/page/EditPost';
import Home from './component/page/Home';
import PreviewPosts from './component/page/PreviewPosts';

function App() {
  return (
    <Router>
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/preview' element={<PreviewPosts />} />
        <Route path='/create' element={<CreatePost />} />
        <Route path='/edit/:id' element={<EditPost />} />
      </Routes>
    </Router>
  );
}

export default App;
