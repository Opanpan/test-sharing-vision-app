import axios from 'axios';
import { useCallback, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import DeleteIcon from '../icon/DeleteIcon';
import EditIcon from '../icon/EditIcon';

export default function Home() {
  const navigate = useNavigate();

  const [posts, setPosts] = useState([]);
  const [filter, setFilter] = useState('publish');

  const getPosts = useCallback(async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/articles/${filter}/10/0`
      );
      setPosts(response?.data);
    } catch (error) {
      console.error('Error fetching posts:', error);
    }
  }, [filter]);

  const onClickDelete = (post) => {
    axios
      .put(`http://localhost:8080/article/${post?.id}`, {
        title: post?.title,
        content: post?.content,
        category: post?.category,
        status: 'thrash',
      })
      .then(() => getPosts())
      .catch((err) => console.log(err));
  };

  const onClickCreate = () => {
    navigate('/create');
  };

  const onClickEdit = (id) => {
    navigate(`/edit/${id}`);
  };

  const onClickTab = (status) => {
    setFilter(status);
  };

  const onClickPreview = () => {
    navigate('/preview');
  };

  useEffect(() => {
    getPosts();
  }, [getPosts]);

  return (
    <div className='p-5'>
      <div className='flex justify-between'>
        <h1 className='mb-5'>All Post</h1>
        <div className='flex gap-5'>
          <button
            className='bg-stone-300 p-2 rounded-md'
            onClick={onClickCreate}
          >
            Create Post
          </button>
          <button
            className='bg-stone-300 p-2 rounded-md'
            onClick={onClickPreview}
          >
            Preview
          </button>
        </div>
      </div>
      <div className='flex gap-x-5 mb-5'>
        <div
          onClick={() => onClickTab('publish')}
          className={`${
            filter === 'publish' ? 'bg-stone-500' : 'bg-stone-300 '
          } p-2 rounded-md hover:cursor-pointer`}
        >
          Published
        </div>
        <div
          onClick={() => onClickTab('draft')}
          className={`${
            filter === 'draft' ? 'bg-stone-500' : 'bg-stone-300 '
          } p-2 rounded-md hover:cursor-pointer`}
        >
          Drafts
        </div>
        <div
          onClick={() => onClickTab('thrash')}
          className={`${
            filter === 'thrash' ? 'bg-stone-500' : 'bg-stone-300 '
          } p-2 rounded-md hover:cursor-pointer`}
        >
          Trashed
        </div>
      </div>
      {posts?.length > 0 && (
        <div>
          <table className='min-w-full bg-white border border-gray-300 rounded-lg shadow-lg'>
            <thead className='bg-stone-300'>
              <tr>
                <th className='px-6 py-3 text-left'>Title</th>
                <th className='px-6 py-3 text-left'>Category</th>
                {filter !== 'thrash' && (
                  <th className='px-6 py-3 text-left'>Action</th>
                )}
              </tr>
            </thead>
            <tbody className='text-gray-700'>
              {posts?.map((post) => {
                return (
                  <>
                    <tr className='border-b'>
                      <td className='px-6 py-4'>{post.title}</td>
                      <td className='px-6 py-4'>{post.category}</td>
                      {filter !== 'thrash' && (
                        <td className='px-6 py-4'>
                          <div className='flex gap-x-3'>
                            <div
                              className='hover:cursor-pointer'
                              onClick={() => onClickDelete(post)}
                            >
                              <DeleteIcon />
                            </div>
                            <div
                              className='hover:cursor-pointer'
                              onClick={() => onClickEdit(post?.id)}
                            >
                              <EditIcon />
                            </div>
                          </div>
                        </td>
                      )}
                    </tr>
                  </>
                );
              })}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
