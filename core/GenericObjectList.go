package core

import (
    "bytes"
    "container/list"
)

type GenericObjectList interface {
    Clone() interface{}
    String() string
}

type GenericObjectListImpl struct {
    list.List

    indentation int
    listName    string // For debugging
    stringRep   string
    separator   string
}

func NewGenericObjectListImpl() *GenericObjectListImpl{
	return &GenericObjectListImpl{};
}

func (this *GenericObjectListImpl) GetIndentation() string {
    var retval bytes.Buffer
    for i := 0; i < this.indentation; i++ {
        retval.WriteString(" ");
    }
    return retval.String();
}

/**
 * Implement the clone method.
 */
func (this *GenericObjectListImpl) Clone() interface{} {
    retval := &GenericObjectListImpl{}

    retval.indentation = this.indentation
    retval.listName = this.listName
    retval.stringRep = this.stringRep
    retval.separator = this.separator

    //gobj.Init();

    return retval
}

/*func (this *GenericObjectListImpl) ConcatenateToTail(objList *GenericObjectListImpl) {
    this.Concatenate(objList, false)
}*/

func (this *GenericObjectListImpl) Concatenate(gol *GenericObjectListImpl){//, topFlag bool) {
    if gol == nil {
        return
    }

    //if !topFlag {
        for e := gol.Front(); e != nil; e = e.Next() {
            this.PushBack(e)
        }
    /*} else {
        // add given items to the end of the list.
        first := this.Front()
        for e := objList.Front(); e != nil; e = e.Next() {
            this.InsertBefore(e, first)
        }
    }*/
}

/**
 * string formatting function.
 */

func (this *GenericObjectListImpl) Sprint(s string) {
    if s == "" {
        this.stringRep += this.GetIndentation()
        this.stringRep += "<null>\n"
        return
    }

    if s == "}" || s == "]" {
        this.indentation--
    }
    this.stringRep += this.GetIndentation()
    this.stringRep += s
    this.stringRep += "\n"
    if s == "{" || s == "[" {
        this.indentation++
    }
}

/**
         * Encode the list in semicolon separated form.
	 * @return an encoded string containing the objects in this list.
         * @since v1.0
*/
func (this *GenericObjectListImpl) String() string {
    if this.Len() == 0 {
        return ""
    }

    var encoding bytes.Buffer //= new StringBuffer();
    for e := this.Front(); e != nil; e = e.Next() {
        if gv, ok := e.Value.(GenericObject); ok {
            encoding.WriteString(gv.String());
        } else {
            encoding.WriteString(e.Value.(string));
        }
    }

    return encoding.String();
}

/**
 *  Set the separator (for encoding the list)
 * @since v1.0
 * @param sep is the new seperator (default is semicolon)
 */
func (this *GenericObjectListImpl) SetSeparator(sep string) {
    this.separator = sep
}
