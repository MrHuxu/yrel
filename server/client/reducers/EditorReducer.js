import immutable from 'immutable';
import {
  CREATE_POST,
  APPEND_POST,
  REPLY_POST,
  REFRESH_POSTS
} from '../actions/EditorActions';

export function editor (state = {
  ids      : immutable.List([]),
  entities : immutable.Map({})
}, action) {
  var copy = Object.assign({}, state);
  const { type, content } = action;

  switch (type) {
    case CREATE_POST:
      copy.ids = copy.ids.push(content.ID);
      copy.entities = copy.entities.set(content.ID, immutable.Map({
        id        : content.ID,
        title     : content.Title,
        content   : content.Content,
        createdAt : new Date(content.CreatedAt),
        appends   : immutable.List([]),
        replies   : immutable.List([])
      }));
      break;

    case APPEND_POST:
      copy.entities = copy.entities.setIn(
        [content.PostID, 'appends'],
        copy.entities.getIn([content.PostID, 'appends']).push({
          text      : content.Text,
          createdAt : new Date(content.CreatedAt)
        })
      );
      break;

    case REPLY_POST:
      copy.entities = copy.entities.setIn(
        [content.PostID, 'replies'],
        copy.entities.getIn([content.PostID, 'replies']).push({
          text      : content.Text,
          replyTo   : content.ReplyTo,
          createdAt : new Date(content.CreatedAt)
        })
      );
      break;

    case REFRESH_POSTS:
      copy.ids = immutable.List(content.map(record => record.ID));
      copy.entities = immutable.Map(content.reduce((prev, cur, index, arr) => {
        prev[cur.ID] = immutable.Map({
          id        : cur.ID,
          title     : cur.Title,
          content   : cur.Content,
          createdAt : new Date(cur.CreatedAt),
          appends   : immutable.List(cur.Appends ? cur.Appends.map(append => {
            return {
              text      : append.Text,
              createdAt : new Date(append.CreatedAt)
            };
          }) : []),
          replies : immutable.List(cur.Replies ? cur.Replies.map(reply => {
            return {
              text      : reply.Text,
              replyTo   : reply.ReplyTo,
              createdAt : new Date(reply.CreatedAt)
            };
          }) : [])
        });

        return prev;
      }, {}));
      break;

    default:
      break;
  }

  return copy;
}
