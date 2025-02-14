import axios from 'axios';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { validateInputForm } from '../../utils';

export default function CreatePost() {
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    title: '',
    content: '',
    category: '',
  });

  const [error, setError] = useState('');

  const handleChange = (e) => {
    setError('');
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (status) => {
    const errorMessage = validateInputForm(formData);

    setError(errorMessage);

    if (errorMessage === '') {
      axios
        .post('http://localhost:8080/article', {
          title: formData.title,
          content: formData.content,
          category: formData.category,
          status: status,
        })
        .then(() => navigate('/'))
        .catch((err) => console.log(err));
    }
  };

  return (
    <div className='max-w-lg mx-auto bg-white p-6 rounded-lg shadow-md mt-10'>
      <h2 className='text-2xl font-bold mb-4'>Create Post</h2>
      {error && <p className='text-red-500 text-sm mb-3'>{error}</p>}

      <div className='mb-5'>
        <label className='block font-medium text-gray-700 mb-3'>Title</label>
        <input
          type='text'
          name='title'
          value={formData.title}
          onChange={handleChange}
          placeholder='Input title...'
          className='w-full px-4 py-2 border rounded-lg focus:ring focus:ring-stone-300 outline-none'
        />
      </div>

      <div className='mb-5'>
        <label className='block font-medium text-gray-700 mb-3'>Category</label>
        <input
          type='text'
          name='category'
          value={formData.category}
          onChange={handleChange}
          placeholder='Input category...'
          className='w-full px-4 py-2 border rounded-lg focus:ring focus:ring-stone-300 outline-none'
        />
      </div>

      <div className='mb-5'>
        <label className='block font-medium text-gray-700 mb-3'>Content</label>
        <textarea
          name='content'
          value={formData.content}
          onChange={handleChange}
          placeholder='Input content...'
          className='w-full px-4 py-2 border rounded-lg focus:ring focus:ring-stone-300 outline-none h-28'
        />
      </div>

      <div className='flex gap-5 mt-5'>
        <button
          onClick={() => handleSubmit('publish')}
          className='w-full bg-stone-500 text-white py-2 rounded-lg hover:bg-stone-600 transition'
        >
          Publish
        </button>

        <button
          onClick={() => handleSubmit('draft')}
          className='w-full bg-stone-500 text-white py-2 rounded-lg hover:bg-stone-600 transition'
        >
          Draft
        </button>
      </div>
    </div>
  );
}
