const matcher = <T extends Record<string, any>>(
  regexp: RegExp,
  fields?: string[],
): ((obj: T) => boolean) => {
  return (obj) => {
    const fields1 = fields || Object.keys(obj as any);
    let found = false;
    fields1.forEach((key) => {
      if (!found) {
        if (typeof obj[key] == "string" && regexp.exec(obj[key])) {
          found = true;
        }
      }
    });
    return found;
  };
};

export function search<T extends Object>(
  collection: T[],
  test: string | string[],
  fields?: string[],
): T[] {
  try {
    let c: T[] = [];
    if (test == undefined || test == null) {
      return c;
    }
    if (typeof test == "string") {
      const regex = new RegExp("\\b" + test, "i");
      c = collection.filter(matcher(regex, fields));
    } else {
      let found: T[] = [];
      test.forEach((t) => {
        const regex = new RegExp("\\b" + t + "\\b", "i");
        found = [...found, ...collection.filter(matcher(regex, fields))];
      });
    }
    return c;
  } catch (error) {
    console.error(error);
    return collection;
  }
}
